package tests

import (
	"os"
	"reflect"
	"testing"
	"time"

	cachetree "soarpipeline/pkg/cachetree"
)

func TestOneXColumnLvmParsesCorrectly(t *testing.T) {
	expected := cachetree.CacheTree{
		PreviewMetadata: cachetree.PreviewMetadata{
			ResultTimestamp: cachetree.TimestampMetadata{
				Date: "2020/08/07",
				Time: "09:47:02.1010842323303222656",
			},
			Operator:     "LaGEsc",
			XColumnNames: []string{"X_Value"},
			YColumnNames: []string{"Pressure", "Temp", "Volume"},
		},
		XColumnNodes: []cachetree.XColumnNode{
			{
				Rows: []float64{0, 0.328878},
			},
		},
		YColumnNodes: []cachetree.YColumnNode{
			{
				Samples:    1,
				Date:       "2020/08/07",
				UnitLabel:  "MPa",
				XDimension: "Time1",
				Rows:       []float64{1.833787, 1.522167},
			},
			{
				Samples:    2,
				Date:       "2020/08/08",
				UnitLabel:  "degC",
				XDimension: "Time2",
				Rows:       []float64{5.479238, 5.310735},
			},
			{
				Samples:    4,
				Date:       "2020/08/09",
				UnitLabel:  "ml",
				XDimension: "Time3",
				Rows:       []float64{0, 89.821400},
			},
		},
	}

	file, err := os.Open("../data/valid_one_x_column.lvm")

	if err != nil {
		t.Errorf("os.Open() error = %v", err)
		return
	}

	defer file.Close()

	parsedCacheTree, err := cachetree.ParseIntoCacheTree(file)

	if err != nil {
		t.Errorf("ParseIntoCacheTree() error = %v", err)
		return
	}

	if !reflect.DeepEqual(parsedCacheTree, expected) {
		t.Errorf("ParseIntoCacheTree() = %v\nwant\n%v", parsedCacheTree, expected)
	}
}

func TestTimeTakenToParseLongLvm(t *testing.T) {
	file, err := os.Open("../data/valid_long.lvm")

	if err != nil {
		t.Errorf("os.Open() error = %v", err)
		return
	}

	defer file.Close()

	start := time.Now()
	_, err = cachetree.ParseIntoCacheTree(file)
	duration := time.Since(start)

	if err != nil {
		t.Errorf("ParseIntoCacheTree() error = %v", err)
		return
	}

	t.Logf("Time taken to parse: %v ms", float64(duration.Microseconds())/1000)
}

func TestMultiXColumnLvmParsesCorrectly(t *testing.T) {
	expected := cachetree.CacheTree{
		PreviewMetadata: cachetree.PreviewMetadata{
			ResultTimestamp: cachetree.TimestampMetadata{
				Date: "2016/08/23",
				Time: "10:45:47.0352557312499836422",
			},
			Operator:     "Ladisk",
			XColumnNames: []string{"(X) Voltage", "(X) Acceleration"},
			YColumnNames: []string{"Voltage", "Acceleration"},
		},
		XColumnNodes: []cachetree.XColumnNode{
			{
				Rows: []float64{0, 1.953125e-5, 3.906250e-5},
			},
			{
				Rows: []float64{0, 1.953125e-5, 3.906250e-5},
			},
		},
		YColumnNodes: []cachetree.YColumnNode{
			{
				Samples:    51200,
				Date:       "2016/08/23",
				UnitLabel:  "Volts",
				XDimension: "Time1",
				Rows:       []float64{-0.035229, -0.034882, -0.034191},
			},
			{
				Samples:    51300,
				Date:       "2016/09/23",
				UnitLabel:  "g",
				XDimension: "Time2",
				Rows:       []float64{0.532608, 0.502991, 0.467541},
			},
		},
	}

	file, err := os.Open("../data/valid_multi_x_column.lvm")

	if err != nil {
		t.Errorf("os.Open() error = %v", err)
		return
	}

	defer file.Close()

	parsedCacheTree, err := cachetree.ParseIntoCacheTree(file)

	if err != nil {
		t.Errorf("ParseIntoCacheTree() error = %v", err)
		return
	}

	if !reflect.DeepEqual(parsedCacheTree, expected) {
		t.Errorf("ParseIntoCacheTree() = %v\nwant\n%v", parsedCacheTree, expected)
	}
}

func TestInvalidLvmFailsToParse(t *testing.T) {
	file, err := os.Open("../data/invalid_1.lvm")

	if err != nil {
		t.Errorf("os.Open() error = %v", err)
		return
	}

	defer file.Close()

	_, err = cachetree.ParseIntoCacheTree(file)

	if err == nil {
		t.Error("ParseIntoCacheTree() error = nil, want error")
	}
}
