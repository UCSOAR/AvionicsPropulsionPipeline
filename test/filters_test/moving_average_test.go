package filters_test

import (
	"reflect"
	"soarpipeline/pkg/filters"
	"testing"
)

// Test for Moving Average Filter.
func TestMovingAvgFilter(t *testing.T) {
	xRows := []float64{1, 2, 3, 4, 5}
	yRows := []float64{10, 20, 30, 40, 50}
	windowSize := 3

	expected := []float64{15, 20, 30, 40, 45}
	result := filters.MovingAvgFilter(xRows, yRows, windowSize)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MovingAvgFilter() = %v, expected %v", result, expected)
	}
}
