package deps

import (
	"fmt"
	"testing"
)

func TestAddFilenameWithPostPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BuildState{}
	filename := "test.txt"
	b.AddFilenameWithPostPrefix(filename)

	if _, ok := b.filenamesWithPostPrefix[filename]; !ok {
		t.Errorf("Expected filename to be added to filenamesWithPostPrefix")
	}
}
