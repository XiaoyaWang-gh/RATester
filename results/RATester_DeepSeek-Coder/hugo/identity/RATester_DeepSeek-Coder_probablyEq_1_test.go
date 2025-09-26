package identity

import (
	"fmt"
	"testing"
)

func TestProbablyEq_1(t *testing.T) {
	type args struct {
		a Identity
		b Identity
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

			if got := probablyEq(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("probablyEq() = %v, want %v", got, tt.want)
			}
		})
	}
}
