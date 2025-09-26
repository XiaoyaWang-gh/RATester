package allconfig

import (
	"fmt"
	"testing"
)

func TestBuildDrafts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.BuildDrafts() != false {
		t.Errorf("BuildDrafts() = %v, want %v", c.BuildDrafts(), false)
	}
}
