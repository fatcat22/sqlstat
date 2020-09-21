package sqlstat

import (
	"fmt"
	"strings"
)

type BooleanAbleStatement interface {
	BooleanAble()
	String() string
}

//
// compare

type compareObject struct {
	col    Column
	symbol string
	value  interface{}
}

func (co *compareObject) BooleanAble() {
	// do nothing
}

func (co *compareObject) String() string {
	return fmt.Sprintf("%s %s '%v'", string(co.col), co.symbol, co.value)
}

func (col Column) Equal(val interface{}) BooleanAbleStatement {
	return &compareObject{
		col:    col,
		symbol: "=",
		value:  val,
	}
}

func (col Column) NotEqual(val interface{}) BooleanAbleStatement {
	return &compareObject{
		col:    col,
		symbol: "!=",
		value:  val,
	}
}

func (col Column) LessOrEqual(val interface{}) BooleanAbleStatement {
	return &compareObject{
		col:    col,
		symbol: "<=",
		value:  val,
	}
}

func (col Column) Great(val interface{}) BooleanAbleStatement {
	return &compareObject{
		col:    col,
		symbol: ">",
		value:  val,
	}
}

//
// logic

type logicObject struct {
	left      BooleanAbleStatement
	rightOnly bool
	logic     string
	right     BooleanAbleStatement
}

func (lo *logicObject) BooleanAble() {
	// do nothing
}

func (lo *logicObject) String() string {
	buf := new(strings.Builder)

	if lo.left != nil {
		buf.WriteString(fmt.Sprintf("(%s)", lo.left.String()))
	}

	if (lo.left != nil) || (lo.left == nil && lo.rightOnly) {
		buf.WriteString(lo.logic)
	}

	buf.WriteString(fmt.Sprintf("(%s)", lo.right.String()))

	return buf.String()
}

func Logic(stat BooleanAbleStatement) *logicObject {
	return &logicObject{
		left: stat,
	}
}

func (lo *logicObject) AND(stat BooleanAbleStatement) BooleanAbleStatement {
	lo.logic = "AND"
	lo.right = stat
	return lo
}

func (lo *logicObject) OR(stat BooleanAbleStatement) BooleanAbleStatement {
	lo.logic = "OR"
	lo.right = stat
	return lo
}

func (lo *logicObject) NOT(stat BooleanAbleStatement) BooleanAbleStatement {
	lo.logic = "NOT"
	lo.right = lo.left
	lo.left = nil
	lo.rightOnly = true
	return lo
}
