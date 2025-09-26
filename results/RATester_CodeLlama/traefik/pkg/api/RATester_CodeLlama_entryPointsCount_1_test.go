package api

import (
	"fmt"
	"testing"
)

func TestEntryPointsCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := udpRouterRepresentation{}
	r.EntryPoints = []string{"127.0.0.1:8080"}
	if r.entryPointsCount() != 1 {
		t.Errorf("Expected 1 entry point, got %d", r.entryPointsCount())
	}
}
