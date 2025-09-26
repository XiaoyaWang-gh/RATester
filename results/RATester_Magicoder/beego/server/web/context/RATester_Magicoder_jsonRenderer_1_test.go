package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestjsonRenderer_1(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want Renderer
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

			if got := jsonRenderer(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonRenderer() = %v, want %v", got, tt.want)
			}
		})
	}
}
