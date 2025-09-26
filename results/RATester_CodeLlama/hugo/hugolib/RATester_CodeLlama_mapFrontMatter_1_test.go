package hugolib

import (
	"fmt"
	"testing"
)

func TestMapFrontMatter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	var rn contentParseInfo
	var source []byte

	// When
	err := rn.mapFrontMatter(source)

	// Then
	if err != nil {
		t.Errorf("mapFrontMatter() error = %v, want nil", err)
		return
	}
}
