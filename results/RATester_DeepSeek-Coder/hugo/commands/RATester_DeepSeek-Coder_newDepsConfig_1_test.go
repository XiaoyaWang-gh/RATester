package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestNewDepsConfig_1(t *testing.T) {
	type args struct {
		conf *commonConfig
	}

	tests := []struct {
		name string
		args args
		want deps.DepsCfg
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

			r := &rootCommand{}
			if got := r.newDepsConfig(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rootCommand.newDepsConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
