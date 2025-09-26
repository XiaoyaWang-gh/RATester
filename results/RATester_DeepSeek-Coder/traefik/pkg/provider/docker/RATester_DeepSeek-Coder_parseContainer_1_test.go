package docker

import (
	"fmt"
	"reflect"
	"testing"

	dockertypes "github.com/docker/docker/api/types"
)

func TestParseContainer_1(t *testing.T) {
	type args struct {
		container dockertypes.ContainerJSON
	}
	tests := []struct {
		name string
		args args
		want dockerData
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

			if got := parseContainer(tt.args.container); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}
