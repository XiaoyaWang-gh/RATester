package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInt_1(t *testing.T) {
	type args struct {
		sourceParam string
		value       string
		dest        interface{}
		bitSize     int
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
			if got := b.int(tt.args.sourceParam, tt.args.value, tt.args.dest, tt.args.bitSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("int() = %v, want %v", got, tt.want)
			}
		})
	}
}
