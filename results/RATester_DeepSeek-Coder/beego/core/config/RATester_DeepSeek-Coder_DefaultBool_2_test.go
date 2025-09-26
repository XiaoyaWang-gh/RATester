package config

import (
	"fmt"
	"testing"
)

func TestDefaultBool_2(t *testing.T) {
	type args struct {
		key        string
		defaultVal bool
	}
	tests := []struct {
		name string
		c    *fakeConfigContainer
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

			if got := tt.c.DefaultBool(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
