package context

import (
	"fmt"
	"testing"
)

func TestIs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		RunMethod: "GET",
	}
	if !input.Is("GET") {
		t.Errorf("TestIs failed")
	}
}
