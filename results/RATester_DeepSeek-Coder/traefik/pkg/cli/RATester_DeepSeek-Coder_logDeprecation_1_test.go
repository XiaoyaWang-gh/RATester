package cli

import (
	"fmt"
	"testing"
)

func TestLogDeprecation_1(t *testing.T) {
	type args struct {
		traefikConfiguration interface{}
		arguments            []string
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

			if got := logDeprecation(tt.args.traefikConfiguration, tt.args.arguments); got != tt.want {
				t.Errorf("logDeprecation() = %v, want %v", got, tt.want)
			}
		})
	}
}
