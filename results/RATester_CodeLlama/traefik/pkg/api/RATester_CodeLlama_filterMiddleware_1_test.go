package api

import (
	"fmt"
	"testing"
)

func TestFilterMiddleware_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &searchCriterion{
		MiddlewareName: "test",
	}

	mns := []string{"test", "test2"}

	if !c.filterMiddleware(mns) {
		t.Errorf("filterMiddleware() = %v, want %v", false, true)
	}
}
