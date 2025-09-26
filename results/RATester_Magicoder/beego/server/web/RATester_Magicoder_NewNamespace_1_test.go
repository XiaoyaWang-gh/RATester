package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewNamespace_1(t *testing.T) {
	type args struct {
		prefix string
		params []LinkNamespace
	}
	tests := []struct {
		name string
		args args
		want *Namespace
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

			if got := NewNamespace(tt.args.prefix, tt.args.params...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
