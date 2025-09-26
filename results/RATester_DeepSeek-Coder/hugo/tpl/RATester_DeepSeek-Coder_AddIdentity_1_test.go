package tpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddIdentity_1(t *testing.T) {
	type args struct {
		t Template
	}
	tests := []struct {
		name string
		args args
		want Template
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

			if got := AddIdentity(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
