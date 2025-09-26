package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeEscaper_1(t *testing.T) {
	type args struct {
		n *nameSpace
	}
	tests := []struct {
		name string
		args args
		want escaper
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

			if got := makeEscaper(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeEscaper() = %v, want %v", got, tt.want)
			}
		})
	}
}
