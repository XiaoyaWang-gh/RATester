package process

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewWithEnvs_1(t *testing.T) {
	type args struct {
		path   string
		params []string
		envs   []string
	}
	tests := []struct {
		name string
		args args
		want *Process
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

			if got := NewWithEnvs(tt.args.path, tt.args.params, tt.args.envs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithEnvs() = %v, want %v", got, tt.want)
			}
		})
	}
}
