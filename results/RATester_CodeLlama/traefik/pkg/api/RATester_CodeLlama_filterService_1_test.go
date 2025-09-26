package api

import (
	"fmt"
	"testing"
)

func TestFilterService_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &searchCriterion{
		ServiceName: "test",
	}

	if !c.filterService("test") {
		t.Error("filterService should return true")
	}

	if c.filterService("test2") {
		t.Error("filterService should return false")
	}
}
