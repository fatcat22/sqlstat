package sqlstat

type Pair struct {
	Column   Column
	Variable interface{}
}

type PairSlice []Pair

func (ps PairSlice) Columns() []Column {
	result := make([]Column, 0, len(ps))
	for _, p := range ps {
		result = append(result, p.Column)
	}
	return result
}

func (ps PairSlice) Variables() []interface{} {
	result := make([]interface{}, 0, len(ps))
	for _, p := range ps {
		result = append(result, p.Variable)
	}
	return result
}

func (ps PairSlice) FormatVariables(f func(v interface{}) string) []string {
	result := make([]string, 0, len(ps))
	for _, p := range ps {
		result = append(result, f(p.Variable))
	}
	return result
}
