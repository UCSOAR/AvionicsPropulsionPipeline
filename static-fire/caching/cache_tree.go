package caching

type TimestampMetadata struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

// Represents metadata for previewing the results of a static fire.
type PreviewMetadata struct {
	ResultTimestamp    TimestampMetadata `json:"resultTimestamp"`
	ProcessedTimestamp TimestampMetadata `json:"processedTimestamp"`
	Operator           string            `json:"operator"`
	XColumnNames       []string          `json:"xColumnNames"`
	YColumnNames       []string          `json:"yColumnNames"`
}

// Represents a X column node in the cache tree.
// Name will be identified by file name.
type XColumnNode struct {
	Dimension string    `json:"dimension"`
	Rows      []float64 `json:"rows"`
}

// Represents a Y column node in the cache tree.
// Name will be identified by file name.
type YColumnNode struct {
	Samples   uint64    `json:"samples"`
	Date      string    `json:"date"`
	UnitLabel string    `json:"unitLabel"`
	Rows      []float64 `json:"rows"`
}

// Represents the file structure to create a cache tree.
type CacheTree struct {
	PreviewMetadata PreviewMetadata
	XColumnNodes    []XColumnNode
	YColumnNodes    []YColumnNode
}
