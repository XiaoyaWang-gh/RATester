package allconfig

import (
	"fmt"
	"testing"
)

func TestIgnoreFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	s := "test"
	if c.IgnoreFile(s) != c.config.C.IgnoreFile(s) {
		t.Errorf("IgnoreFile() = %v, want %v", c.IgnoreFile(s), c.config.C.IgnoreFile(s))
	}
}
