package metrics

import (
	"fmt"
	"testing"
)

func TestIsSvcEnabled_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registry := &standardRegistry{
		svcEnabled: true,
	}

	if !registry.IsSvcEnabled() {
		t.Error("Expected svcEnabled to be true, but it was false")
	}
}
