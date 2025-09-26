package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func Testsort_2(t *testing.T) {
	type args struct {
		p pairList
	}
	tests := []struct {
		name string
		args args
		want any
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

			if got := tt.args.p.sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
