package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUint16_1(t *testing.T) {
	type args struct {
		sourceParam string
		dest        *uint16
	}
	tests := []struct {
		name string
		args args
		want *ValueBinder
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

			b := &ValueBinder{}
			if got := b.Uint16(tt.args.sourceParam, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}
