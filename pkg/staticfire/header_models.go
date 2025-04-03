package staticfire

import "fmt"

// Enumerates the possible values for the `MultiHeadings` field in an LVM file.
type MultiHeadingsValue uint8

const (
	MultiHeadingsUnknown MultiHeadingsValue = iota
	MultiHeadingsYes     MultiHeadingsValue = iota
	MultiHeadingsNo      MultiHeadingsValue = iota
)

// Enumerates the possible values for the `XColumns` field in an LVM file.
type XColumnsValue uint8

const (
	XColumnsUnknown XColumnsValue = iota
	XColumnsOne     XColumnsValue = iota
	XColumnsMulti   XColumnsValue = iota
)

// Enumerates the possible field separators in an LVM file.
type Fieldseparator rune

const (
	FieldseparatorUnknown Fieldseparator = 0
	FieldseparatorTab     Fieldseparator = '\t'
	FieldseparatorSpace   Fieldseparator = ' '
)

// Represents a parsed LVM entry header.
type ParsedEntryHeader struct {
	Separator     Fieldseparator     `json:"separator"`
	MultiHeadings MultiHeadingsValue `json:"multiHeadings"`
	XColumns      XColumnsValue      `json:"xColumns"`
	Operator      string             `json:"operator"`
	Date          string             `json:"date"`
	Time          string             `json:"time"`
}

// Represents a parsed LVM channel header.
// The data stored is relevant to the purposes of this project.
// It is guaranteed that all arrays will have the same length as `ChannelCount`.
type ParsedChannelHeader struct {
	ChannelCount int       `json:"channelCount"`
	Samples      []int     `json:"samples"`
	Dates        []string  `json:"dates"`
	Times        []string  `json:"times"`
	YUnitLabels  []string  `json:"yUnitLabels"`
	XDimensions  []string  `json:"xDimensions"`
	InitialXs    []float64 `json:"initialXs"`
	DeltaXs      []float64 `json:"deltaXs"`
}

func ParseMultiHeadingsValue(multiHeadingsText string) (MultiHeadingsValue, error) {
	switch multiHeadingsText {
	case "Yes":
		return MultiHeadingsYes, nil
	case "No":
		return MultiHeadingsNo, nil
	default:
		return MultiHeadingsUnknown, fmt.Errorf("invalid multi headings value: %s", multiHeadingsText)
	}
}

func (m MultiHeadingsValue) String() string {
	switch m {
	case 1:
		return "Yes"
	case 2:
		return "No"
	default:
		return "Unknown"
	}
}

func ParseXColumnsValue(xColumnsText string) (XColumnsValue, error) {
	switch xColumnsText {
	case "One":
		return XColumnsOne, nil
	case "Multi":
		return XColumnsMulti, nil
	default:
		return XColumnsUnknown, fmt.Errorf("invalid X columns value: %s", xColumnsText)
	}
}

func (x XColumnsValue) String() string {
	switch x {
	case 1:
		return "One"
	case 2:
		return "Multi"
	default:
		return "Unknown"
	}
}

func ParseFieldseparator(separatorText string) (Fieldseparator, error) {
	switch separatorText {
	case "Tab":
		return FieldseparatorTab, nil
	case "Space":
		return FieldseparatorTab, nil
	default:
		return FieldseparatorUnknown, fmt.Errorf("invalid separator: %s", separatorText)
	}
}
