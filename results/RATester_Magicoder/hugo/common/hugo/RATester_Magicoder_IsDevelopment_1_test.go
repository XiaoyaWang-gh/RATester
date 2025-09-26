package hugo

import (
	"fmt"
	"testing"
)

func TestIsDevelopment_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	info := HugoInfo{
		Environment: "development",
	}

	if !info.IsDevelopment() {
		t.Error("Expected IsDevelopment to return true")
	}

	info.Environment = "production"

	if info.IsDevelopment() {
		t.Error("Expected IsDevelopment to return false")
	}
}
