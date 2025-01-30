package storage

import (
	parser "example.com/static-fire/parser"
)

type ResultMetadata struct {
	Operator      string   `json:"operator"`
	ResultDate    string   `json:"resultDate"`
	ResultTime    string   `json:"resultTime"`
	ProcessedDate string   `json:"processedDate"`
	ProcessedTime string   `json:"processedTime"`
	XColumnNames  []string `json:"xColumnNames"`
	YColumnNames  []string `json:"yColumnNames"`
}

type HeaderMetadata struct {
	EntryHeader   parser.ParsedEntryHeader
	ChannelHeader parser.ParsedChannelHeader
}
