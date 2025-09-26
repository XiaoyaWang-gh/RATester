package version

import (
	"fmt"
	"testing"
)

func TestFull_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := "version"
	if got := Full(); got != expected {
		t.Errorf("Full() = %v, want %v", got, expected)
	}
}
