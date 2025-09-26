package metrics

import (
	"context"
	"fmt"
	"testing"
)

func TestRegisterPromState_1(t *testing.T) {
	type args struct {
		ctx context.Context
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

			if got := registerPromState(tt.args.ctx); got != tt.want {
				t.Errorf("registerPromState() = %v, want %v", got, tt.want)
			}
		})
	}
}
