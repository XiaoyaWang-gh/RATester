package config

import (
	"fmt"
	"testing"
)

func TestDefaultString_4(t *testing.T) {
	type args struct {
		key        string
		defaultVal string
	}
	tests := []struct {
		name string
		c    *BaseConfiger
		args args
		want string
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

			if got := tt.c.DefaultString(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultString() = %v, want %v", got, tt.want)
			}
		})
	}
}
