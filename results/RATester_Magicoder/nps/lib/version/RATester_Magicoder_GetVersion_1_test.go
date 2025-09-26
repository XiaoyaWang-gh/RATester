package version

import (
	"fmt"
	"testing"
)

func TestGetVersion_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := "0.26.0"
	result := GetVersion()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
