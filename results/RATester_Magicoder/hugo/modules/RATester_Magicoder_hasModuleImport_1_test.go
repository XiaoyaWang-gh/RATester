package modules

import (
	"fmt"
	"testing"
)

func TesthasModuleImport_1(t *testing.T) {
	type args struct {
		c Config
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

			if got := tt.args.c.hasModuleImport(); got != tt.want {
				t.Errorf("hasModuleImport() = %v, want %v", got, tt.want)
			}
		})
	}
}
