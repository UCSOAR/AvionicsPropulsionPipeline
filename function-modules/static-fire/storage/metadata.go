package storage

import (
	parser "example.com/static-fire/parser"
)

type ResultMetadata struct {
	Operator     string   `json:"operator"`
	Date         string   `json:"date"`
	Time         string   `json:"time"`
	XColumnNames []string `json:"x_column_names"`
	YColumnNames []string `json:"y_column_names"`
}

type ChannelMetadata struct {
	Header parser.ParsedChannelHeader `json:"data"`
}
