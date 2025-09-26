package hugolib

import (
	"fmt"
	"testing"
)

func TestInitRenderHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	pco := &pageContentOutput{}

	// When
	err := pco.initRenderHooks()

	// Then
	if err != nil {
		t.Errorf("initRenderHooks() unexpected error = %v", err)
		return
	}
}
