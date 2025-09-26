package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMustUint8_1(t *testing.T) {
	type args struct {
		sourceParam string
		dest        *uint8
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
			if got := b.MustUint8(tt.args.sourceParam, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}
