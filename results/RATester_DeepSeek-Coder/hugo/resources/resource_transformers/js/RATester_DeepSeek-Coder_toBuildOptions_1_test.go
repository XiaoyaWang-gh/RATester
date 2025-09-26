package js

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/evanw/esbuild/pkg/api"
)

func TestToBuildOptions_1(t *testing.T) {
	type args struct {
		opts Options
	}
	tests := []struct {
		name             string
		args             args
		wantBuildOptions api.BuildOptions
		wantErr          bool
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

			gotBuildOptions, err := toBuildOptions(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("toBuildOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBuildOptions, tt.wantBuildOptions) {
				t.Errorf("toBuildOptions() = %v, want %v", gotBuildOptions, tt.wantBuildOptions)
			}
		})
	}
}
