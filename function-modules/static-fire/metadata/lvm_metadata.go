package metadata

type LvmMetadata struct {
	ProcessedTimestamp string   `json:"processed_timestamp"`
	Operator           string   `json:"operator"`
	ColumnNames        []string `json:"column_names"`
}
