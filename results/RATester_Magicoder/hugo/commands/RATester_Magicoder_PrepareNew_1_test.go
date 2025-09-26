package commands

import (
	"fmt"
	"testing"
)

func TestPrepareNew_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	detector := &fileChangeDetector{
		current: map[string]uint64{
			"file1": 123,
			"file2": 456,
		},
		prev: map[string]uint64{
			"file3": 789,
			"file4": 012,
		},
	}

	detector.PrepareNew()

	if len(detector.prev) != 2 || len(detector.current) != 0 {
		t.Errorf("PrepareNew() failed. Expected prev map to be empty and current map to have 2 elements, but got %d and %d", len(detector.prev), len(detector.current))
	}

	if detector.prev["file1"] != 123 || detector.prev["file2"] != 456 {
		t.Errorf("PrepareNew() failed. Expected prev map to have file1 and file2 with checksum 123 and 456, but got %d and %d", detector.prev["file1"], detector.prev["file2"])
	}
}
