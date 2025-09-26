package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxSize_1(t *testing.T) {
	type args struct {
		obj interface{}
		max int
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
			if got := v.MaxSize(tt.args.obj, tt.args.max, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
