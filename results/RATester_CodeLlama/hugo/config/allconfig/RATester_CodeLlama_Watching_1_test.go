package allconfig

import (
	"fmt"
	"testing"
)

func TestWatching_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.Watching() {
		t.Error("Watching() should be false")
	}
}
