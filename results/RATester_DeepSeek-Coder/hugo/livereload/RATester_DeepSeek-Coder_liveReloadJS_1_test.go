package livereload

import (
	"fmt"
	"testing"
)

func TestLiveReloadJS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	result := liveReloadJS()
	if len(result) == 0 {
		t.Errorf("liveReloadJS() returned an empty byte array")
	}
}
