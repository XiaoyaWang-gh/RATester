package hugo

import (
	"fmt"
	"testing"
)

func TestIsProduction_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	info := HugoInfo{
		Environment: "production",
	}

	if !info.IsProduction() {
		t.Error("Expected IsProduction to return true")
	}

	info.Environment = "development"

	if info.IsProduction() {
		t.Error("Expected IsProduction to return false")
	}
}
