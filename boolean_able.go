package sqlstat

import (
	"fmt"
	"strings"
)

type booleanAbleStatement interface {
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

func (col Column) Equal(val interface{}) booleanAbleStatement {
	return &compareObject{
		col:    col,
		symbol: "=",
		value:  val,
	}
}

func (col Column) NotEqual(val interface{}) booleanAbleStatement {
	return &compareObject{
		col:    col,
		symbol: "!=",
		value:  val,
	}
}

func (col Column) LessOrEqual(val interface{}) booleanAbleStatement {
	return &compareObject{
		col:    col,
		symbol: "<=",
		value:  val,
	}
}

func (col Column) Great(val interface{}) booleanAbleStatement {
	return &compareObject{
		col:    col,
		symbol: ">",
		value:  val,
	}
}

//
// logic

type logicObject struct {
	left  booleanAbleStatement
	logic string
	right booleanAbleStatement
}

func (lo *logicObject) BooleanAble() {
	// do nothing
}

func (lo *logicObject) String() string {
	buf := new(strings.Builder)

	if lo.left != nil {
		buf.WriteString(fmt.Sprintf("(%s)", lo.left.String()))
	}
	buf.WriteString(lo.logic)
	buf.WriteString(fmt.Sprintf("(%s)", lo.right.String()))

	return buf.String()
}

func Logic(stat booleanAbleStatement) *logicObject {
	return &logicObject{
		left: stat,
	}
}

func (lo *logicObject) AND(stat booleanAbleStatement) booleanAbleStatement {
	lo.logic = "AND"
	lo.right = stat
	return lo
}

func (lo *logicObject) OR(stat booleanAbleStatement) booleanAbleStatement {
	lo.logic = "OR"
	lo.right = stat
	return lo
}

func (lo *logicObject) NOT(stat booleanAbleStatement) booleanAbleStatement {
	lo.logic = "NOT"
	lo.right = lo.left
	lo.left = nil
	return lo
}
