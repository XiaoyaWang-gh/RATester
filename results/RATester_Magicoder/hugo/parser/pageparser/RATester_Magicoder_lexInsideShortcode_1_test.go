package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestlexInsideShortcode_1(t *testing.T) {
	tests := []struct {
		name string
		l    *pageLexer
		want stateFunc
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

			if got := lexInsideShortcode(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexInsideShortcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
