package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinSize_1(t *testing.T) {
	type args struct {
		obj interface{}
		min int
		key string
	}
	tests := []struct {
		name string
		args args
		want *Result
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

			v := &Validation{}
			if got := v.MinSize(tt.args.obj, tt.args.min, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
