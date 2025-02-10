package cachetree

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
}

// Represents a X column node in the cache tree.
// Name will be identified by file name.
type XColumnNode struct {
	Rows []float64 `json:"rows"`
}

// Represents a Y column node in the cache tree.
// Name will be identified by file name.
type YColumnNode struct {
	Samples    int       `json:"samples"`
	Date       string    `json:"date"`
	UnitLabel  string    `json:"unitLabel"`
	XDimension string    `json:"xDimension"`
	Rows       []float64 `json:"rows"`
}

// Represents the file structure to create a cache tree.
type CacheTree struct {
	PreviewMetadata PreviewMetadata
	XColumnNodes    []XColumnNode
	YColumnNodes    []YColumnNode
}
