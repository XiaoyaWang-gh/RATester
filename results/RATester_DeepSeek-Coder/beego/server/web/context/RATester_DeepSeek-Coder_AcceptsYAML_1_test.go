package context

import (
	"fmt"
	"testing"
)

func TestAcceptsYAML_1(t *testing.T) {
	type args struct {
		input *BeegoInput
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

			if got := tt.args.input.AcceptsYAML(); got != tt.want {
				t.Errorf("BeegoInput.AcceptsYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}
