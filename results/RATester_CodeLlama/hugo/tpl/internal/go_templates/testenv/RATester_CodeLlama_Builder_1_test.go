package testenv

import (
	"fmt"
	"testing"
)

func TestBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if Builder() != "golang" {
		t.Errorf("Builder() = %v, want %v", Builder(), "golang")
	}
}
