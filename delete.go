package sqlstat

import (
	"fmt"
	"strings"
)

type deleteObject struct {
	tableName string
	where     BooleanAbleStatement
}

func DeleteFrom(tableName string) *deleteObject {
	return &deleteObject{
		tableName: tableName,
	}
}

func (do *deleteObject) String() string {
	buf := new(strings.Builder)

	buf.WriteString(fmt.Sprintf("DELETE FROM %s", do.tableName))

	if do.where != nil {
		buf.WriteString(fmt.Sprintf(" WHERE %s", do.where.String()))
	}

	return buf.String()
}

func (do *deleteObject) Where(where BooleanAbleStatement) *deleteObject {
	do.where = where
	return do
}
