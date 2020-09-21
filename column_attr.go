package sqlstat

import (
	"fmt"
	"strconv"
	"strings"
)

type ColumnAttr struct {
	name      string
	ctype     string
	isNotNULL bool

	hasDefault   bool
	defaultValue string

	isAutoIncrement bool
}

func (c Column) Attr() *ColumnAttr {
	return &ColumnAttr{
		name: string(c),
	}
}

func (ca *ColumnAttr) String() string {
	fields := make([]string, 0, 5)

	fields = append(fields, ca.name)
	fields = append(fields, ca.ctype)
	if ca.isNotNULL {
		fields = append(fields, "NOT NULL")
	}
	if ca.hasDefault {
		fields = append(fields, "DEFAULT "+ca.defaultValue)
	}
	if ca.isAutoIncrement {
		fields = append(fields, "AUTO_INCREMENT")
	}

	return strings.Join(fields, " ")
}

func (ca *ColumnAttr) Type(t string) *ColumnAttr {
	ca.ctype = t
	return ca
}

func (ca *ColumnAttr) NotNULL() *ColumnAttr {
	ca.isNotNULL = true
	return ca
}

func (ca *ColumnAttr) DefaultEmptyString() *ColumnAttr {
	ca.hasDefault = true
	ca.defaultValue = "''"
	return ca
}

func (ca *ColumnAttr) DefaultString(s string) *ColumnAttr {
	ca.hasDefault = true
	ca.defaultValue = fmt.Sprintf("'%s'", s)
	return ca
}

func (ca *ColumnAttr) DefaultInt(v int) *ColumnAttr {
	ca.hasDefault = true
	ca.defaultValue = strconv.Itoa(v)
	return ca
}

func (ca *ColumnAttr) AutoIncrement() *ColumnAttr {
	ca.isAutoIncrement = true
	return ca
}

func (ca *ColumnAttr) GetColumnName() string {
	return ca.name
}
