package livereload

import (
	"fmt"
	"testing"
)

func TestliveReloadJS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	result := liveReloadJS()
	if len(result) == 0 {
		t.Error("Expected non-empty result, but got empty")
	}
}
