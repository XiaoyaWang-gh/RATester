package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestNewFrom_2(t *testing.T) {
	type args struct {
		params maps.Params
	}
	tests := []struct {
		name string
		args args
		want Provider
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

			if got := NewFrom(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
