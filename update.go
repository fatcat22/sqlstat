package sqlstat

import (
	"fmt"
	"strings"
)

type updateObject struct {
	tableName string

	pairs PairSlice
	where BooleanAbleStatement
}

func Update(tableName string) *updateObject {
	return &updateObject{
		tableName: tableName,
		pairs:     make(PairSlice, 0),
	}
}

func (uo *updateObject) String() string {
	buf := new(strings.Builder)
	buf.WriteString(fmt.Sprintf("UPDATE %s SET ", uo.tableName))

	var sets []string
	for _, p := range uo.pairs {
		sets = append(sets, fmt.Sprintf("%s = '%v'", p.Column, p.Variable))
	}
	buf.WriteString(strings.Join(sets, ","))

	if uo.where != nil {
		buf.WriteString(fmt.Sprintf(" WHERE %s", uo.where.String()))
	}

	return buf.String()
}

func (uo *updateObject) With(col Column, val interface{}) *updateObject {
	uo.pairs = append(uo.pairs, Pair{Column: col, Variable: val})
	return uo
}

func (uo *updateObject) WithPairSlice(ps PairSlice) *updateObject {
	uo.pairs = append(uo.pairs, ps...)
	return uo
}

func (uo *updateObject) Where(w BooleanAbleStatement) *updateObject {
	uo.where = w
	return uo
}
