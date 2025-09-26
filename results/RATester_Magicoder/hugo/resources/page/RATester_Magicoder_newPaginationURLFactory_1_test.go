package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewPaginationURLFactory_1(t *testing.T) {
	type args struct {
		d TargetPathDescriptor
	}
	tests := []struct {
		name string
		args args
		want paginationURLFactory
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

			if got := newPaginationURLFactory(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPaginationURLFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
