package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMustInt32s_1(t *testing.T) {
	type args struct {
		sourceParam string
		dest        *[]int32
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
			if got := b.MustInt32s(tt.args.sourceParam, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}
