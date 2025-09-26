package template

import (
	"fmt"
	"testing"
)

func TestIsInTag_1(t *testing.T) {
	tests := []struct {
		name string
		s    state
		want bool
	}{
		{"stateTag", stateTag, true},
		{"stateAttrName", stateAttrName, true},
		{"stateAfterName", stateAfterName, true},
		{"stateBeforeValue", stateBeforeValue, true},
		{"stateAttr", stateAttr, true},
		{"stateUnknown", 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isInTag(tt.s); got != tt.want {
				t.Errorf("isInTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
