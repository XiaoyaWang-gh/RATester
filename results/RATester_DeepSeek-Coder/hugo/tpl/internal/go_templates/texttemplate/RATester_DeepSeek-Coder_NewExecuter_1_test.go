package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewExecuter_1(t *testing.T) {
	type args struct {
		helper ExecHelper
	}
	tests := []struct {
		name string
		args args
		want Executer
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

			if got := NewExecuter(tt.args.helper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExecuter() = %v, want %v", got, tt.want)
			}
		})
	}
}
