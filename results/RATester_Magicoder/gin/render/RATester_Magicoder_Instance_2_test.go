package render

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInstance_2(t *testing.T) {
	type args struct {
		name string
		data any
	}
	tests := []struct {
		name string
		r    HTMLProduction
		args args
		want Render
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

			if got := tt.r.Instance(tt.args.name, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTMLProduction.Instance() = %v, want %v", got, tt.want)
			}
		})
	}
}
