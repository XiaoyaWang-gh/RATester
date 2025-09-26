package config

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_4(t *testing.T) {
	type args struct {
		key        string
		defaultVal float64
	}
	tests := []struct {
		name string
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

			if got := DefaultFloat(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
