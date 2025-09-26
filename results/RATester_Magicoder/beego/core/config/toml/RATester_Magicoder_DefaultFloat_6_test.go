package toml

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_6(t *testing.T) {
	type args struct {
		key        string
		defaultVal float64
	}
	tests := []struct {
		name string
		c    *configContainer
		args args
		want float64
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

			if got := tt.c.DefaultFloat(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("configContainer.DefaultFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
