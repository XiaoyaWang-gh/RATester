package utils

import (
	"fmt"
	"testing"
)

func TestNewBeeMap_1(t *testing.T) {
	t.Run("TestNewBeeMap", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		bm := NewBeeMap()
		if bm == nil {
			t.Errorf("NewBeeMap() = %v, want %v", bm, "non-nil")
		}
	})
}
