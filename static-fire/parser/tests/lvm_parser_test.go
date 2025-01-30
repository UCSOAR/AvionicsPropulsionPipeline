package tests

import (
	"os"
	"reflect"
	"testing"

	parser "github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/parser"
)

func TestValidLVMParsesCorrectly(t *testing.T) {
	seperator, err := parser.ParseFieldSeperator("Tab")

	if err != nil {
		t.Errorf("ParseFieldSeperator() error = %v", err)
		return
	}

	multiHeadings, err := parser.ParseMultiHeadingsValue("No")

	if err != nil {
		t.Errorf("ParseMultiHeadingsValue() error = %v", err)
		return
	}

	xColumns, err := parser.ParseXColumnsValue("One")

	if err != nil {
		t.Errorf("ParseXColumnsValue() error = %v", err)
	}

	expected := parser.ParsedLvm{
		Headers: parser.HeaderMetadata{
			EntryHeader: parser.ParsedEntryHeader{
				Seperator:     seperator,
				MultiHeadings: multiHeadings,
				XColumns:      xColumns,
				Operator:      "RothneyPC",
				Date:          "2023/09/21",
				Time:          "14:41:45.5111323470585990021",
			},
			ChannelHeader: parser.ParsedChannelHeader{
				ChannelCount: 7,
				Samples:      []uint64{100, 100, 100, 100, 100, 100, 100},
				Dates:        []string{"2023/09/21", "2023/09/21", "2023/09/21", "2023/09/21", "2023/09/21", "2023/09/21", "2023/09/21"},
				Times:        []string{"14:41:45.5111323470585990021", "14:41:45.5111463466995406219", "14:41:45.5111603463404822418", "14:41:45.5111323470585990021", "14:41:45.5111743459814238617", "14:41:45.5111463466995406219", "14:41:45.5111883456223654816"},
				YUnitLabels:  []string{"Some", "PSI", "kg", "N", "PSI", "Volts", "PSI"},
				XDimensions:  []string{"Time", "Time", "Time", "Time", "Time", "Time", "Time"},
				InitialXs:    []float64{0, 0, 0, 0, 0, 0, 0},
				DeltaXs:      []float64{0.001, 0.001, 0.001, 0.001, 0.001, 0.001, 0.001},
			},
		},
		SvData: parser.ParsedSv{
			ColumnCount: 8,
			ColumnNames: []string{"X_Value", "Injector Temp", "NOX Pressure", "Ox Tank Load Cell", "Thrust Load Cell", "Chamber Pressure", "Light Sensor", "Nitrous Bottle"},
			Data: [][]float64{
				{0, -250.069013, 0.568622, -20.072851, 9.56483, 0, 0.040771, 1.886051},
				{0.001, -250.069013, 0.568622, -20.063231, 10.154943, 0, 0.020755, 0},
			},
		},
	}

	file, err := os.Open("./files/valid.lvm")

	if err != nil {
		t.Errorf("os.Open() error = %v", err)
		return
	}

	parsedLvm, err := parser.ParseLvm(file)

	if err != nil {
		t.Errorf("ParseMainLvm() error = %v", err)
		return
	}

	if !reflect.DeepEqual(parsedLvm, expected) {
		t.Errorf("ParseMainLvm() EntryHeader = %v, want %v", parsedLvm, expected)
	}
}
