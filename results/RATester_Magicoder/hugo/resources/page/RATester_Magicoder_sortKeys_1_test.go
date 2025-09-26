package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsortKeys_1(t *testing.T) {
	type args struct {
		examplePage Page
		v           []reflect.Value
		order       string
	}
	tests := []struct {
		name string
		args args
		want []reflect.Value
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

			if got := sortKeys(tt.args.examplePage, tt.args.v, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
