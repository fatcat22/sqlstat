package sqlstat

import "strings"

type Column string

func NewColumn(name string) Column {
	return Column(name)
}

func JoinColumn(cols []Column, sep string) string {
	if len(cols) == 0 {
		return ""
	}

	buf := new(strings.Builder)
	buf.WriteString(string(cols[0]))
	for _, s := range cols[1:] {
		buf.WriteString(sep)
		buf.WriteString(string(s))
	}

	return buf.String()

}
