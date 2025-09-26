package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/simplecobra"
)

func TestNewNewCommand_1(t *testing.T) {
	type fields struct {
		rootCmd  *rootCommand
		commands []simplecobra.Commander
	}
	tests := []struct {
		name   string
		fields fields
		want   *newCommand
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

			got := newNewCommand()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNewCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
