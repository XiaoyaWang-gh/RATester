package commands

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHugoBuilder_1(t *testing.T) {
	type args struct {
		r              *rootCommand
		s              *serverCommand
		onConfigLoaded []func(reloaded bool) error
	}
	tests := []struct {
		name string
		args args
		want *hugoBuilder
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

			if got := newHugoBuilder(tt.args.r, tt.args.s, tt.args.onConfigLoaded...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHugoBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
