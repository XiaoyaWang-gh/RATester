package task

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewMapSorter_1(t *testing.T) {
	type args struct {
		m map[string]Tasker
	}
	tests := []struct {
		name string
		args args
		want *MapSorter
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

			if got := NewMapSorter(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMapSorter() = %v, want %v", got, tt.want)
			}
		})
	}
}
