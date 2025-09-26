package web

import (
	"fmt"
	"testing"
)

func TestLockViewPaths_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lockViewPaths()
	if !beeViewPathTemplateLocked {
		t.Errorf("Expected beeViewPathTemplateLocked to be true, but got false")
	}
}
