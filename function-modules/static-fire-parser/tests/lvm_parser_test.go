package staticFireParserTests

import (
	"reflect"
	"testing"

	staticFireParser "example.com/static-fire-parser"
)

func TestValidLVMParsesCorrectly(t *testing.T) {
	rawLvmText := `LabVIEW Measurement	
Writer_Version	2
Reader_Version	2
Separator	Tab
Decimal_Separator	.
Multi_Headings	No
X_Columns	One
Time_Pref	Absolute
Operator	RothneyPC
Date	2023/09/21
Time	14:41:45.5111323470585990021
***End_of_Header***	
	
Channels	7							
Samples	100	100	100	100	100	100	100	
Date	2023/09/21	2023/09/21	2023/09/21	2023/09/21	2023/09/21	2023/09/21	2023/09/21	
Time	14:41:45.5111323470585990021	14:41:45.5111463466995406219	14:41:45.5111603463404822418	14:41:45.5111323470585990021	14:41:45.5111743459814238617	14:41:45.5111463466995406219	14:41:45.5111883456223654816	
Y_Unit_Label	Some	PSI	kg	N	PSI	Volts	PSI	
X_Dimension	Time	Time	Time	Time	Time	Time	Time	
X0	0.0000000000000000E+0	0.0000000000000000E+0	0.0000000000000000E+0	0.0000000000000000E+0	0.0000000000000000E+0	0.0000000000000000E+0	0.0000000000000000E+0	
Delta_X	0.001000	0.001000	0.001000	0.001000	0.001000	0.001000	0.001000	
***End_of_Header***								
X_Value	Injector Temp	NOX Pressure	Ox Tank Load Cell	Thrust Load Cell	Chamber Pressure	Light Sensor	Nitrous Bottle	Comment
0.000000	-250.069013	0.568622	-20.072851	9.564830	0.000000	0.040771	1.886051	Standby
0.001000	-250.069013	0.568622	-20.063231	10.154943	0.000000	0.020755	0.000000`

	expected := staticFireParser.ParsedLvm{
		EntryHeader: staticFireParser.ParsedEntryHeader{
			Seperator:     '\t',
			HasOneXColumn: true,
			Operator:      "RothneyPC",
			Date:          "2023/09/21",
			Time:          "14:41:45.5111323470585990021",
		},
		ChannelHeader: staticFireParser.ParsedChannelHeader{
			ChannelCount: 7,
			Samples:      []uint64{100, 100, 100, 100, 100, 100, 100},
			Dates:        []string{"2023/09/21", "2023/09/21", "2023/09/21", "2023/09/21", "2023/09/21", "2023/09/21", "2023/09/21"},
			Times:        []string{"14:41:45.5111323470585990021", "14:41:45.5111463466995406219", "14:41:45.5111603463404822418", "14:41:45.5111323470585990021", "14:41:45.5111743459814238617", "14:41:45.5111463466995406219", "14:41:45.5111883456223654816"},
			YunitLabels:  []string{"Some", "PSI", "kg", "N", "PSI", "Volts", "PSI"},
			Xdimensions:  []string{"Time", "Time", "Time", "Time", "Time", "Time", "Time"},
			InitialXs:    []float64{0, 0, 0, 0, 0, 0, 0},
			DeltaXs:      []float64{0.001, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001},
		},
		SvData: staticFireParser.ParsedSv{
			ColumnCount: 8,
			ColumnNames: []string{"X_Value", "Injector Temp", "NOX Pressure", "Ox Tank Load Cell", "Thrust Load Cell", "Chamber Pressure", "Light Sensor", "Nitrous Bottle", "Comment"},
			Data: [][]float64{
				{0, -250.069013, 0.568622, -20.072851, 9.56483, 0, 0.040771, 1.886051},
				{0.001, -250.069013, 0.568622, -20.063231, 10.154943, 0, 0.020755, 0},
			},
		},
	}

	parsedLvm, err := staticFireParser.ParseMainLvm(rawLvmText)

	if err != nil {
		t.Errorf("ParseMainLvm() error = %v", err)
		return
	}

	if !reflect.DeepEqual(parsedLvm.EntryHeader, expected.EntryHeader) {
		t.Errorf("ParseMainLvm() EntryHeader = %v, want %v", parsedLvm.EntryHeader, expected.EntryHeader)
	}
}
