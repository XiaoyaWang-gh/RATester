package render

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInstance_1(t *testing.T) {
	type args struct {
		name string
		data any
	}
	tests := []struct {
		name string
		r    HTMLDebug
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
				t.Errorf("HTMLDebug.Instance() = %v, want %v", got, tt.want)
			}
		})
	}
}
