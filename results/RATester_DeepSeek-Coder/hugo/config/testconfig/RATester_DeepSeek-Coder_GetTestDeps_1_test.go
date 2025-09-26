package testconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/deps"
	"github.com/spf13/afero"
)

func TestGetTestDeps_1(t *testing.T) {
	type args struct {
		fs   afero.Fs
		cfg  config.Provider
		init []func(*deps.Deps)
	}
	tests := []struct {
		name string
		args args
		want *deps.Deps
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

			if got := GetTestDeps(tt.args.fs, tt.args.cfg, tt.args.init...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTestDeps() = %v, want %v", got, tt.want)
			}
		})
	}
}
