package test

import (
	"fmt"
	"testing"
)

func TestAssetNames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := []string{"file1.txt", "file2.txt", "file3.txt"}
	actual := AssetNames()

	if len(actual) != len(expected) {
		t.Errorf("Expected length of %d, but got %d", len(expected), len(actual))
	}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Expected %s at index %d, but got %s", expected[i], i, v)
		}
	}
}
