package hugolib

import (
	"fmt"
	"testing"
)

func TestshouldRender_1(t *testing.T) {
	type args struct {
		cfg *BuildCfg
		p   *pageState
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

			if got := tt.args.cfg.shouldRender(tt.args.p); got != tt.want {
				t.Errorf("shouldRender() = %v, want %v", got, tt.want)
			}
		})
	}
}
