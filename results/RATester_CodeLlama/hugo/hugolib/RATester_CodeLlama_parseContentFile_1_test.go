package hugolib

import (
	"fmt"
	"testing"
)

func TestParseContentFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &cachedContent{}
	source := []byte("")
	if err := c.parseContentFile(source); err != nil {
		t.Errorf("parseContentFile() error = %v", err)
		return
	}
}
