package set_test

import (
	"soarpipeline/pkg/set"
	"testing"
)

func TestAddAndHasInSet(t *testing.T) {
	const testStr = "test"

	set := set.NewHashSet[string]()
	set.Put(testStr)

	if !set.Has(testStr) {
		t.Errorf("Expected value to exist in the set")
	}

	if set.Has("nonexistent") {
		t.Errorf("Expected value to not exist in the set")
	}
}

func TestRemoveFromSet(t *testing.T) {
	const testInt = 42

	set := set.NewHashSet[int]()
	set.Put(testInt)

	if !set.Has(testInt) {
		t.Errorf("Expected value to exist in the set")
	}

	set.Remove(testInt)

	if set.Has(testInt) {
		t.Errorf("Expected value to be removed from the set")
	}
}
