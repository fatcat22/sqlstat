package sqlstat

import (
	"fmt"
	"strings"
)

type bitwiseObject struct {
	column         Column
	singleOperator bool
	bitSymbol      string
	bitValue       interface{}

	logicSymbol string
	logicValue  interface{}
}

func (c Column) BitAnd(v interface{}) *bitwiseObject {
	return &bitwiseObject{
		column:    c,
		bitSymbol: "&",
		bitValue:  v,
	}
}

func (bo *bitwiseObject) BooleanAble() {
	// do nothing
}

func (bo *bitwiseObject) String() string {
	if len(bo.bitSymbol) == 0 {
		return ""
	}

	buf := new(strings.Builder)
	if bo.singleOperator {
		buf.WriteString(fmt.Sprintf("%s%s", bo.bitSymbol, bo.column))
	} else {
		buf.WriteString(fmt.Sprintf("%s%s%v", bo.column, bo.bitSymbol, bo.bitValue))
	}

	if len(bo.logicSymbol) != 0 {
		old := buf.String()
		buf.Reset()
		buf.WriteString(fmt.Sprintf("(%s) %s '%v'", old, bo.logicSymbol, bo.logicValue))
	}

	return buf.String()
}

func (bo *bitwiseObject) Equal(v interface{}) booleanAbleStatement {
	bo.logicSymbol = "="
	bo.logicValue = v
	return bo
}
