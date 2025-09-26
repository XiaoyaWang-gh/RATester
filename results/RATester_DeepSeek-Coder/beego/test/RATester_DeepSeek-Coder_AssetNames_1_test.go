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

	expected := []string{
		"asset1.txt",
		"asset2.txt",
		"asset3.txt",
		// add more expected asset names here
	}

	actual := AssetNames()

	if len(expected) != len(actual) {
		t.Errorf("Expected %d asset names, got %d", len(expected), len(actual))
		return
	}

	for i, name := range actual {
		if name != expected[i] {
			t.Errorf("Expected asset name %d to be %s, got %s", i, expected[i], name)
		}
	}
}
