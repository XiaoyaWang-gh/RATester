package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconsumeToSpace_1(t *testing.T) {
	tests := []struct {
		name string
		l    *pageLexer
		want []byte
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

			tt.l.consumeToSpace()
			if !reflect.DeepEqual(tt.l.current(), tt.want) {
				t.Errorf("consumeToSpace() = %v, want %v", tt.l.current(), tt.want)
			}
		})
	}
}
