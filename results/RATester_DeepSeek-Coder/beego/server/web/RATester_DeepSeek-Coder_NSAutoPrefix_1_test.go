package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNSAutoPrefix_1(t *testing.T) {
	type args struct {
		prefix string
		c      ControllerInterface
	}
	tests := []struct {
		name string
		args args
		want LinkNamespace
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

			if got := NSAutoPrefix(tt.args.prefix, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NSAutoPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
