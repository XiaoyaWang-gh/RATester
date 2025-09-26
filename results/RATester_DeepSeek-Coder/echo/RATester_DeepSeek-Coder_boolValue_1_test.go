package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBoolValue_1(t *testing.T) {
	type args struct {
		sourceParam    string
		dest           *bool
		valueMustExist bool
	}
	tests := []struct {
		name string
		b    *ValueBinder
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

			if got := tt.b.boolValue(tt.args.sourceParam, tt.args.dest, tt.args.valueMustExist); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boolValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
