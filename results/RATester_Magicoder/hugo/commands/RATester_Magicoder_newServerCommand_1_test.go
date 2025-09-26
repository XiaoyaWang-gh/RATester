package commands

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewServerCommand_1(t *testing.T) {
	type args struct {
		uninstall bool
	}
	tests := []struct {
		name string
		args args
		want *serverCommand
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

			if got := newServerCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newServerCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
