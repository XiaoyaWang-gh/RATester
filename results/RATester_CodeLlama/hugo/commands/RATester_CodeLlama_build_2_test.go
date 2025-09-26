package commands

import (
	"fmt"
	"testing"
)

func TestBuild_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &hugoBuilder{}
	err := c.build()
	if err != nil {
		t.Errorf("build() failed with error: %s", err)
	}
}
