package config

import (
	"fmt"
	"testing"
)

func TestGetBool_1(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name string
		c    *defaultConfigProvider
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

			if got := tt.c.GetBool(tt.args.k); got != tt.want {
				t.Errorf("defaultConfigProvider.GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
