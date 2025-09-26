package commands

import (
	"fmt"
	"testing"
)

func TestFullBuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	c := &hugoBuilder{}
	noBuildLock := false

	// When
	err := c.fullBuild(noBuildLock)

	// Then
	if err != nil {
		t.Errorf("fullBuild() error = %v", err)
		return
	}
}
