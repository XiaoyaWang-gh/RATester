package config

import (
	"fmt"
	"testing"
)

func TestDefaultInt64_4(t *testing.T) {
	type args struct {
		key        string
		defaultVal int64
	}
	tests := []struct {
		name string
		c    *BaseConfiger
		args args
		want int64
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

			if got := tt.c.DefaultInt64(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
