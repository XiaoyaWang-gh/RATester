package web

import (
	"fmt"
	"testing"
)

func TestlockViewPaths_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lockViewPaths()
	if !beeViewPathTemplateLocked {
		t.Error("Expected beeViewPathTemplateLocked to be true, but it was false")
	}
}
