package sqlstat

import (
	"fmt"
	"strings"
)

type orderType string

const (
	DESC = orderType("DESC")
	ASC  = orderType("ASC")
)

type orderInfo struct {
	column Column
	otype  orderType
}

type selectObject struct {
	tableName string

	isCount bool
	columns []Column

	// where
	where BooleanAbleStatement

	// order by
	orderBy []orderInfo

	//limit
	hasLimit  bool
	hasOffset bool
	offset    int
	count     int
}

func Select(cols ...Column) *selectObject {
	return &selectObject{
		columns: cols,
	}
}

func SelectCount(cols ...Column) *selectObject {
	return &selectObject{
		isCount: true,
		columns: cols,
	}
}

func (so selectObject) String() string {
	buf := new(strings.Builder)

	columnStrs := JoinColumn(so.columns, ",")
	switch {
	case so.isCount:
		columnStrs = fmt.Sprintf("COUNT(%s)", columnStrs)
	}
	buf.WriteString(fmt.Sprintf("SELECT %s FROM %s", columnStrs, so.tableName))

	if so.where != nil {
		buf.WriteString(fmt.Sprintf(" WHERE %s", so.where.String()))
	}

	if len(so.orderBy) != 0 {
		var ordersInfo []string
		for _, o := range so.orderBy {
			ordersInfo = append(ordersInfo, fmt.Sprintf("%s %s", o.column, o.otype))
		}
		buf.WriteString(fmt.Sprintf(" ORDER BY %s", strings.Join(ordersInfo, ",")))
	}

	if so.hasLimit {
		if so.hasOffset {
			buf.WriteString(fmt.Sprintf(" LIMIT %d, %d", so.offset, so.count))
		} else {
			buf.WriteString(fmt.Sprintf(" LIMIT %d", so.count))
		}
	}

	return buf.String()
}

func (so *selectObject) From(tableName string) *selectObject {
	so.tableName = tableName
	return so
}

func (so *selectObject) Where(where BooleanAbleStatement) *selectObject {
	so.where = where
	return so
}

func (so *selectObject) OrderBy(col Column, otype orderType) *selectObject {
	so.orderBy = append(so.orderBy, orderInfo{col, otype})
	return so
}

func (so *selectObject) Limit(offset, count int) *selectObject {
	so.hasLimit = true
	so.hasOffset = true
	so.offset = offset
	so.count = count
	return so
}

func (so *selectObject) LimitCount(count int) *selectObject {
	so.hasLimit = true
	so.hasOffset = false
	so.count = count
	return so
}
