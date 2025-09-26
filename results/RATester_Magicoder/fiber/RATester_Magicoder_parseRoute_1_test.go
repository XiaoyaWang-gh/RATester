package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestparseRoute_1(t *testing.T) {
	type args struct {
		pattern          string
		customConstraint CustomConstraint
	}
	tests := []struct {
		name string
		args args
		want routeParser
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

			if got := parseRoute(tt.args.pattern, tt.args.customConstraint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
