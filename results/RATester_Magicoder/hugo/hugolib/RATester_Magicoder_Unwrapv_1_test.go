package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnwrapv_1(t *testing.T) {
	type fields struct {
		weight0   int
		pageState *pageState
	}
	tests := []struct {
		name   string
		fields fields
		want   any
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

			p := pageWithWeight0{
				weight0:   tt.fields.weight0,
				pageState: tt.fields.pageState,
			}
			if got := p.Unwrapv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unwrapv() = %v, want %v", got, tt.want)
			}
		})
	}
}
