package source

import (
	"fmt"
	"testing"
)

func TestIsContentAdapter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &File{}
	if fi.IsContentAdapter() {
		t.Error("IsContentAdapter() should be false")
	}
}
