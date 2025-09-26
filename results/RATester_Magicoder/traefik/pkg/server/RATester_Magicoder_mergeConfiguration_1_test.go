package server

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestMergeConfiguration_1(t *testing.T) {
	type args struct {
		configurations     dynamic.Configurations
		defaultEntryPoints []string
	}
	tests := []struct {
		name string
		args args
		want dynamic.Configuration
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

			if got := mergeConfiguration(tt.args.configurations, tt.args.defaultEntryPoints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}
