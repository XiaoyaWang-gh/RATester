package hugolib

import (
	"fmt"
	"testing"
)

func TestInitCommonProviders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var pp pagePaths
	var ps *pageState

	// act
	err := ps.initCommonProviders(pp)

	// assert
	if err != nil {
		t.Errorf("Did not expect error, but got %s", err)
	}
}
