package hugolib

import (
	"fmt"
	"testing"
)

func TestEq_5(t *testing.T) {
	type args struct {
		other any
	}
	tests := []struct {
		name string
		p    *pageState
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

			if got := tt.p.Eq(tt.args.other); got != tt.want {
				t.Errorf("pageState.Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}
