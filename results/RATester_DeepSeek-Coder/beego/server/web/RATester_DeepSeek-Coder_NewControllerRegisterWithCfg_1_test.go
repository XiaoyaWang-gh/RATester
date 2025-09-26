package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewControllerRegisterWithCfg_1(t *testing.T) {
	type args struct {
		cfg *Config
	}
	tests := []struct {
		name string
		args args
		want *ControllerRegister
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

			if got := NewControllerRegisterWithCfg(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewControllerRegisterWithCfg() = %v, want %v", got, tt.want)
			}
		})
	}
}
