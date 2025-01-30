package tests

import (
	"reflect"
	"testing"

	parser "example.com/static-fire/parser"
)

func TestValidChannelHeaderTextParsesCorrectly(t *testing.T) {
	rawHeaderText := `Channels 2
Samples	51200		51200	
Date	2016/08/23		2016/08/23	
Time	10:45:47.0352557312499836422		10:45:47.0352557312499836422	
Y_Unit_Label	Volts		g	
X_Dimension	Time		Time	
X0	0.0000000000000000E+0		0.0000000000000000E+0	
Delta_X	1.953125E-5		1.953125E-5	
`

	expected := parser.ParsedChannelHeader{
		ChannelCount: 2,
		Samples:      []uint64{51200, 51200},
		Dates:        []string{"2016/08/23", "2016/08/23"},
		Times:        []string{"10:45:47.0352557312499836422", "10:45:47.0352557312499836422"},
		YUnitLabels:  []string{"Volts", "g"},
		XDimensions:  []string{"Time", "Time"},
		InitialXs:    []float64{0.0000000000000000, 0.0000000000000000},
		DeltaXs:      []float64{1.953125e-5, 1.953125e-5},
	}

	result, err := parser.ParseChannelHeader(rawHeaderText)

	if err != nil {
		t.Errorf("ParseChannelHeader() error = %v", err)
		return
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseChannelHeader() = %v, want %v", result, expected)
	}
}

func TestMismatchedChannelCountReturnsError(t *testing.T) {
	rawHeaderText := `Channels 2
Samples	51200
Date	2016/08/23
Time	10:45:47.0352557312499836422
Y_Unit_Label	Volts
X_Dimension	Time
X0	0.0000000000000000E+0
Delta_X	1.953125E-5
`

	_, err := parser.ParseChannelHeader(rawHeaderText)

	if err == nil {
		t.Errorf("ParseChannelHeader() error = %v, want an error", err)
	}
}

func TestIncompleteChannelHeaderReturnsError(t *testing.T) {
	rawHeaderText := `Channels 2
Samples	51200		51200
Date	2016/08/23		2016/08/23
Time	10:45:47.0352557312499836422		10:45:47.0352557312499836422
Y_Unit_Label	Volts		g
X_Dimension	Time		Time
X0	0.0000000000000000E+0		0.0000000000000000E+0
`

	_, err := parser.ParseChannelHeader(rawHeaderText)

	if err == nil {
		t.Errorf("ParseChannelHeader() error = %v, want an error", err)
	}
}
