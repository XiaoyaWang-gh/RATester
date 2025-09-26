package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/simplecobra"
)

func TestnewListCommand_1(t *testing.T) {
	type args struct {
		cd *simplecobra.Commandeer
		r  *rootCommand
	}
	tests := []struct {
		name string
		args args
		want *listCommand
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

			if got := newListCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newListCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
