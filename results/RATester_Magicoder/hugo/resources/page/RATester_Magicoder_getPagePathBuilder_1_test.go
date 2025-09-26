package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetPagePathBuilder_1(t *testing.T) {
	type args struct {
		d TargetPathDescriptor
	}
	tests := []struct {
		name string
		args args
		want *pagePathBuilder
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

			if got := getPagePathBuilder(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPagePathBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
