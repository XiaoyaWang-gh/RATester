package fiber

import (
	"fmt"
	"testing"
)

func TestMatch_1(t *testing.T) {
	type args struct {
		detectionPath string
		path          string
		params        *[30]string
	}
	tests := []struct {
		name string
		r    *Route
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

			if got := tt.r.match(tt.args.detectionPath, tt.args.path, tt.args.params); got != tt.want {
				t.Errorf("Route.match() = %v, want %v", got, tt.want)
			}
		})
	}
}
