package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniqueRouteStack_1(t *testing.T) {
	type args struct {
		stack []*Route
	}
	tests := []struct {
		name string
		args args
		want []*Route
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

			if got := uniqueRouteStack(tt.args.stack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniqueRouteStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
