package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMustUint32s_1(t *testing.T) {
	type args struct {
		sourceParam string
		dest        *[]uint32
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

			if got := (&ValueBinder{}).MustUint32s(tt.args.sourceParam, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustUint32s() = %v, want %v", got, tt.want)
			}
		})
	}
}
