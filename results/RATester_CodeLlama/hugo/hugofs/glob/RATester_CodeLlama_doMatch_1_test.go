package glob

import (
	"fmt"
	"testing"
)

func TestDoMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &FilenameFilter{}
	filename := "filename"
	isDir := false
	if got := f.doMatch(filename, isDir); got != true {
		t.Errorf("doMatch() = %v, want %v", got, true)
	}
}
