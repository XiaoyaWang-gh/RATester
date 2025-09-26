package tplimpl

import (
	"fmt"
	"testing"
)

func TestNewTemplateState_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	// CONTEXT_1
	// CONTEXT_2
	// CONTEXT_3
	// CONTEXT_4
	// CONTEXT_5
	// CONTEXT_6
	// CONTEXT_7
	// CONTEXT_8
	// CONTEXT_9
	// METHOD UNDER TEST
	// ASSERTIONS
}
