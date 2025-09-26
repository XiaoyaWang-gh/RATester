package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewPageFinder_1(t *testing.T) {
	type args struct {
		m *pageMap
	}
	tests := []struct {
		name string
		args args
		want *pageFinder
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

			if got := newPageFinder(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPageFinder() = %v, want %v", got, tt.want)
			}
		})
	}
}
