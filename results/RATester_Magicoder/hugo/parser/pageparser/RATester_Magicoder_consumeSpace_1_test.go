package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconsumeSpace_1(t *testing.T) {
	tests := []struct {
		name string
		l    *pageLexer
		want *pageLexer
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

			tt.l.consumeSpace()
			if !reflect.DeepEqual(tt.l, tt.want) {
				t.Errorf("consumeSpace() = %v, want %v", tt.l, tt.want)
			}
		})
	}
}
