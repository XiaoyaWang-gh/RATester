package framework

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetEnvs_1(t *testing.T) {
	type args struct {
		envs []string
	}
	tests := []struct {
		name string
		f    *Framework
		args args
		want []string
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

			tt.f.SetEnvs(tt.args.envs)
			if !reflect.DeepEqual(tt.f.osEnvs, tt.want) {
				t.Errorf("got %v, want %v", tt.f.osEnvs, tt.want)
			}
		})
	}
}
