package docker

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/docker/docker/client"
)

func TestNewDynConfBuilder_1(t *testing.T) {
	type args struct {
		configuration Shared
		apiClient     client.APIClient
	}
	tests := []struct {
		name string
		args args
		want *DynConfBuilder
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

			if got := NewDynConfBuilder(tt.args.configuration, tt.args.apiClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDynConfBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
