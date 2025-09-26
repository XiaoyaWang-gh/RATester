package server

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestApplyModel_1(t *testing.T) {
	type args struct {
		cfg dynamic.Configuration
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

			if got := applyModel(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
