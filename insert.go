package sqlstat

import (
	"fmt"
	"strings"
)

type insertObject struct {
	tableName   string
	pairs       PairSlice
	updatePairs PairSlice
}

func InsertInto(tableName string) *insertObject {
	return &insertObject{
		tableName:   tableName,
		pairs:       make([]Pair, 0),
		updatePairs: make([]Pair, 0),
	}
}

func (ins *insertObject) String() string {
	if len(ins.pairs) == 0 {
		return ""
	}

	columns := ins.pairs.Columns()
	values := ins.pairs.FormatVariables(func(v interface{}) string { return fmt.Sprintf("'%v'", v) })

	buf := new(strings.Builder)
	buf.WriteString(fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)",
		ins.tableName, JoinColumn(columns, ","), strings.Join(values, ",")))

	if len(ins.updatePairs) != 0 {
		var sets []string
		for _, p := range ins.updatePairs {
			sets = append(sets, fmt.Sprintf("%s = '%v'", p.Column, p.Variable))
		}
		buf.WriteString(fmt.Sprintf(" ON DUPLICATE KEY UPDATE %s", strings.Join(sets, ",")))
	}

	return buf.String()
}

func (ins *insertObject) With(col Column, val interface{}) *insertObject {
	ins.pairs = append(ins.pairs, Pair{col, val})
	return ins
}

func (ins *insertObject) WithPairSlice(ps PairSlice) *insertObject {
	ins.pairs = append(ins.pairs, ps...)
	return ins
}

func (ins *insertObject) UpdateWith(col Column, val interface{}) *insertObject {
	ins.updatePairs = append(ins.updatePairs, Pair{col, val})
	return ins
}

func (ins *insertObject) UpdateWithPairSlice(ps PairSlice) *insertObject {
	ins.updatePairs = append(ins.updatePairs, ps...)
	return ins
}
