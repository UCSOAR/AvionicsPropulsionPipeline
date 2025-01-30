package caching

// Represents metadata for previewing the results of a static fire.
type PreviewMetadata struct {
	Operator      string   `json:"operator"`
	ResultDate    string   `json:"resultDate"`
	ResultTime    string   `json:"resultTime"`
	ProcessedDate string   `json:"processedDate"`
	ProcessedTime string   `json:"processedTime"`
	XColumns      []string `json:"xColumns"` // Names of the X columns
	YColumns      []string `json:"yColumns"` // Names of the Y columns
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
