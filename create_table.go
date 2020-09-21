package sqlstat

import (
	"fmt"
	"strings"
)

type CreateTableObject struct {
	tableName  string
	columns    []*ColumnAttr
	index      []Column
	uniqueKey  []Column
	primaryKey []Column

	engineName     string
	defaultCharset string
	autoIncrement  int
}

func CreateTable(name string) *CreateTableObject {
	return &CreateTableObject{
		tableName: name,
		columns:   make([]*ColumnAttr, 0),

		engineName:     "InnoDB",
		defaultCharset: "utf8",
		autoIncrement:  1,
	}
}

func (cto *CreateTableObject) String() string {
	fields := make([]string, 0)

	colAttrStrs := make([]string, 0, len(cto.columns))
	for _, ca := range cto.columns {
		colAttrStrs = append(colAttrStrs, ca.String())
	}
	if len(colAttrStrs) != 0 {
		fields = append(fields, strings.Join(colAttrStrs, ","))
	}

	if len(cto.index) != 0 {
		fields = append(fields, fmt.Sprintf("INDEX (%s)", JoinColumn(cto.index, ",")))
	}
	if len(cto.uniqueKey) != 0 {
		fields = append(fields, fmt.Sprintf("UNIQUE KEY(%s)", JoinColumn(cto.uniqueKey, ",")))
	}
	if len(cto.primaryKey) != 0 {
		fields = append(fields, fmt.Sprintf("PRIMARY KEY (%s)", JoinColumn(cto.primaryKey, ",")))
	}

	return fmt.Sprintf(`CREATE TABLE %s(%s) ENGINE=%s  DEFAULT CHARSET=%s AUTO_INCREMENT = %d`,
		cto.tableName, strings.Join(fields, ","),
		cto.engineName, cto.defaultCharset, cto.autoIncrement)
}

func (cto *CreateTableObject) AppendColumn(cattr *ColumnAttr) *CreateTableObject {
	cto.columns = append(cto.columns, cattr)
	return cto
}

func (cto *CreateTableObject) Index(cols []Column) *CreateTableObject {
	cto.index = append(cto.index, cols...)
	return cto
}

func (cto *CreateTableObject) UniqueKey(cols []Column) *CreateTableObject {
	cto.uniqueKey = append(cto.uniqueKey, cols...)
	return cto
}

func (cto *CreateTableObject) PrimaryKey(cols []Column) *CreateTableObject {
	cto.primaryKey = append(cto.primaryKey, cols...)
	return cto
}

func (cto *CreateTableObject) Engine(name string) *CreateTableObject {
	cto.engineName = name
	return cto
}

func (cto *CreateTableObject) DefaultCharset(c string) *CreateTableObject {
	cto.defaultCharset = c
	return cto
}

func (cto *CreateTableObject) AutoIncrementBegin(i int) *CreateTableObject {
	cto.autoIncrement = i
	return cto
}
