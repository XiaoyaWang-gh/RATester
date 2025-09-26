package helpers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestProcessingStatsTable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var w bytes.Buffer
	var stats []*ProcessingStats

	// when
	ProcessingStatsTable(&w, stats...)

	// then
	// ...
}
