package models

import (
	"fmt"
	"testing"
)

func TestEmpty_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{}
	if !mc.Empty() {
		t.Errorf("mc.Empty() = %v, want %v", mc.Empty(), true)
	}
}
