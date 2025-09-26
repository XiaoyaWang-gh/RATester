package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewIntegrationTestBuilder_1(t *testing.T) {
	type args struct {
		conf IntegrationTestConfig
	}
	tests := []struct {
		name string
		args args
		want *IntegrationTestBuilder
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

			if got := NewIntegrationTestBuilder(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntegrationTestBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
