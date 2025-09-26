package media

import (
	"fmt"
	"testing"
)

func TesthasSuffix_2(t *testing.T) {
	type args struct {
		suffix string
	}
	tests := []struct {
		name string
		m    Type
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

			if got := tt.m.hasSuffix(tt.args.suffix); got != tt.want {
				t.Errorf("hasSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
