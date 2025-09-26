package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShape_2(t *testing.T) {
	type args struct {
		d int
		v int
	}
	tests := []struct {
		name string
		t    pageTrees
		args args
		want *pageTrees
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

			if got := tt.t.Shape(tt.args.d, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pageTrees.Shape() = %v, want %v", got, tt.want)
			}
		})
	}
}
