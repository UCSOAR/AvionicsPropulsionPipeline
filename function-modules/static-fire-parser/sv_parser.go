package staticFireParser

type ParsedSv struct {
	ColumnCount uint64
	ColumnNames []string
	Data        [][]float64 // Row major order
}

func ParseSv(rawSvText string, delimiter rune) (ParsedSv, error) {
	sv := ParsedSv{}

	return sv, nil
}
