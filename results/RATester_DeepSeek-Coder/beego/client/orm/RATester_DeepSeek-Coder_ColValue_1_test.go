package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestColValue_1(t *testing.T) {
	type args struct {
		opt   operator
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
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

			if got := ColValue(tt.args.opt, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
