package hugolib

import (
	"fmt"
	"testing"
)

func TestWeight0_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	page := pageWithWeight0{weight0: 10}
	result := page.Weight0()
	if result != 10 {
		t.Errorf("Expected Weight0 to be 10, but got %d", result)
	}
}
