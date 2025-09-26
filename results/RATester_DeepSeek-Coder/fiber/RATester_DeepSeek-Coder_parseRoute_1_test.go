package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseRoute_1(t *testing.T) {
	type args struct {
		pattern           string
		customConstraints []CustomConstraint
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

			if got := parseRoute(tt.args.pattern, tt.args.customConstraints...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
