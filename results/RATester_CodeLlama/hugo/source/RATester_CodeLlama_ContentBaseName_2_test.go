package source

import (
	"fmt"
	"testing"
)

func TestContentBaseName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &File{}
	if got := fi.ContentBaseName(); got != "" {
		t.Errorf("ContentBaseName() = %v, want %v", got, "")
	}
}
