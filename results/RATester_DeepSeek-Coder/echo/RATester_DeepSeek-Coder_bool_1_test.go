package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBool_1(t *testing.T) {
	type args struct {
		sourceParam string
		value       string
		dest        *bool
	}
	tests := []struct {
		name string
		args args
		want *bool
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
			if got := b.bool(tt.args.sourceParam, tt.args.value, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
