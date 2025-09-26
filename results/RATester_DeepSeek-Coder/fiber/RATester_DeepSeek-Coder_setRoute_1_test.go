package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetRoute_1(t *testing.T) {
	type args struct {
		route *Route
	}
	tests := []struct {
		name string
		c    *DefaultCtx
		args args
		want *Route
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

			tt.c.setRoute(tt.args.route)
			if !reflect.DeepEqual(tt.c.route, tt.want) {
				t.Errorf("DefaultCtx.setRoute() = %v, want %v", tt.c.route, tt.want)
			}
		})
	}
}
