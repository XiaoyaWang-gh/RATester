package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewPageRef_1(t *testing.T) {
	type args struct {
		p *pageState
	}
	tests := []struct {
		name string
		args args
		want pageRef
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

			if got := newPageRef(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPageRef() = %v, want %v", got, tt.want)
			}
		})
	}
}
