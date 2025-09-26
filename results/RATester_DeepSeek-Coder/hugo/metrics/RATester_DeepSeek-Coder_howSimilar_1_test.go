package metrics

import (
	"fmt"
	"testing"
)

func TestHowSimilar_1(t *testing.T) {
	type args struct {
		a any
		b any
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

			if got := howSimilar(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("howSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
