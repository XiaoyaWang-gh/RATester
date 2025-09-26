package source

import (
	"fmt"
	"testing"
)

func TestSection_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var fi *File
	if got := fi.Section(); got != "" {
		t.Errorf("File.Section() = %v, want \"\"", got)
	}
}
