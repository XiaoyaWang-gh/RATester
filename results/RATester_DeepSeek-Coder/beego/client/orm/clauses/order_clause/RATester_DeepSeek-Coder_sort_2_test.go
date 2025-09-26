package order_clause

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort_2(t *testing.T) {
	type args struct {
		sort Sort
	}
	tests := []struct {
		name string
		args args
		want Option
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

			if got := sort(tt.args.sort); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
