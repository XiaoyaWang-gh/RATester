package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestJoinRange_1(t *testing.T) {
	tests := []struct {
		name string
		c0   context
		rc   *rangeContext
		want context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := joinRange(tt.c0, tt.rc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("joinRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
