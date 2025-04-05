package staticfire

type TimestampMetadata struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

// Represents metadata for previewing the results of a static fire.
type PreviewMetadata struct {
	ResultTimestamp TimestampMetadata `json:"resultTimestamp"`
	Operator        string            `json:"operator"`
	XColumnNames    []string          `json:"xColumnNames"`
	YColumnNames    []string          `json:"yColumnNames"`
	TotalRows       int               `json:"totalRows"`
}

// Represents a column node in the cache tree.
type ColumnNode struct {
	Rows []float64 `json:"rows"`
}

// Represents metadata for the Y column in the cache tree.
type YColumnMetadata struct {
	Samples    int    `json:"samples"`
	Date       string `json:"date"`
	UnitLabel  string `json:"unitLabel"`
	XDimension string `json:"xDimension"`
}

// Represents the file structure to create a cache tree.
type CacheTree struct {
	PreviewMetadata PreviewMetadata
	YColumnMetadata []YColumnMetadata // Matches the length of YRows
	XColumnNodes    []ColumnNode
	YColumnNodes    []ColumnNode
}
