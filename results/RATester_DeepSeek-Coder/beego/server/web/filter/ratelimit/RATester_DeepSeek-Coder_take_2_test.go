package ratelimit

import (
	"fmt"
	"testing"
)

func TestTake_2(t *testing.T) {
	type args struct {
		amount uint
	}
	tests := []struct {
		name string
		b    *tokenBucket
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

			if got := tt.b.take(tt.args.amount); got != tt.want {
				t.Errorf("tokenBucket.take() = %v, want %v", got, tt.want)
			}
		})
	}
}
