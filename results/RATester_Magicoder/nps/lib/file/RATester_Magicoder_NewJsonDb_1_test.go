package file

import (
	"fmt"
	"testing"
)

func TestNewJsonDb_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	runPath := "/path/to/run"
	db := NewJsonDb(runPath)

	if db.RunPath != runPath {
		t.Errorf("Expected RunPath to be %s, but got %s", runPath, db.RunPath)
	}

	// Add more test cases as needed
}
