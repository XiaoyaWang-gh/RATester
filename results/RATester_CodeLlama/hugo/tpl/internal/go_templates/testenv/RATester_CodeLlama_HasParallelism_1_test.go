package testenv

import (
	"fmt"
	"testing"
)

func TestHasParallelism_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if HasParallelism() {
		t.Log("HasParallelism() is true")
	} else {
		t.Error("HasParallelism() is false")
	}
}
