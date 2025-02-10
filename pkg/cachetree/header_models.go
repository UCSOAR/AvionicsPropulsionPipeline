package cachetree

import "fmt"

type MultiHeadingsValue uint8

const (
	MultiHeadingsUnknown MultiHeadingsValue = 0
	MultiHeadingsYes     MultiHeadingsValue = 1
	MultiHeadingsNo      MultiHeadingsValue = 2
)

type XColumnsValue uint8

const (
	XColumnsUnknown XColumnsValue = 0
	XColumnsOne     XColumnsValue = 1
	XColumnsMulti   XColumnsValue = 2
)

type FieldSeperator rune

const (
	FieldSeperatorUnknown FieldSeperator = 0
	FieldSeperatorTab     FieldSeperator = '\t'
	FieldSeperatorSpace   FieldSeperator = ' '
)

type ParsedEntryHeader struct {
	Seperator     FieldSeperator     `json:"seperator"`
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
		return MultiHeadingsUnknown, fmt.Errorf("Invalid multi headings value: %s", multiHeadingsText)
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
		return XColumnsUnknown, fmt.Errorf("Invalid X columns value: %s", xColumnsText)
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

func ParseFieldSeperator(seperatorText string) (FieldSeperator, error) {
	switch seperatorText {
	case "Tab":
		return FieldSeperatorTab, nil
	case "Space":
		return FieldSeperatorTab, nil
	default:
		return FieldSeperatorUnknown, fmt.Errorf("Invalid seperator: %s", seperatorText)
	}
}
