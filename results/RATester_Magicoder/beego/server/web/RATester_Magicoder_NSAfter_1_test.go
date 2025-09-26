package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNSAfter_1(t *testing.T) {
	type args struct {
		filterList []FilterFunc
	}
	tests := []struct {
		name string
		args args
		want LinkNamespace
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

			if got := NSAfter(tt.args.filterList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NSAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}
