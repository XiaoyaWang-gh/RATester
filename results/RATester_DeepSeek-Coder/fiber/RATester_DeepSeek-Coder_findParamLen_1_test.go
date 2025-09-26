package fiber

import (
	"fmt"
	"testing"
)

func TestFindParamLen_1(t *testing.T) {
	type args struct {
		s       string
		segment *routeSegment
	}
	tests := []struct {
		name string
		args args
		want int
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

			if got := findParamLen(tt.args.s, tt.args.segment); got != tt.want {
				t.Errorf("findParamLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
