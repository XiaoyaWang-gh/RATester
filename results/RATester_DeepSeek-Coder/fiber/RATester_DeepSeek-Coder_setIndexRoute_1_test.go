package fiber

import (
	"fmt"
	"testing"
)

func TestSetIndexRoute_1(t *testing.T) {
	type args struct {
		route int
	}
	tests := []struct {
		name string
		c    *DefaultCtx
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

			tt.c.setIndexRoute(tt.args.route)
			if got := tt.c.getIndexRoute(); got != tt.want {
				t.Errorf("DefaultCtx.setIndexRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
