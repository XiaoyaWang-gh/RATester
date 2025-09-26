package diagrams

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoat_1(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		d    *Namespace
		args args
		want SVGDiagram
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

			if got := tt.d.Goat(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Goat() = %v, want %v", got, tt.want)
			}
		})
	}
}
