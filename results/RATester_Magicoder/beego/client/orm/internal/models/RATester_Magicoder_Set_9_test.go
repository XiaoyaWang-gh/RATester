package models

import (
	"fmt"
	"testing"
)

func TestSet_9(t *testing.T) {
	testCases := []struct {
		name string
		d    uint16
		want uint16
	}{
		{"TestSet1", 1, 1},
		{"TestSet2", 65535, 65535},
		{"TestSet3", 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e := PositiveSmallIntegerField(0)
			e.Set(tc.d)
			if e.Value() != tc.want {
				t.Errorf("Expected %d, got %d", tc.want, e.Value())
			}
		})
	}
}
