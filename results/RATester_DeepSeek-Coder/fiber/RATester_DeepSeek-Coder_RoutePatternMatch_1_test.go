package fiber

import (
	"fmt"
	"testing"
)

func TestRoutePatternMatch_1(t *testing.T) {
	type args struct {
		path    string
		pattern string
		cfg     Config
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := RoutePatternMatch(tt.args.path, tt.args.pattern, tt.args.cfg); got != tt.want {
				t.Errorf("RoutePatternMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
