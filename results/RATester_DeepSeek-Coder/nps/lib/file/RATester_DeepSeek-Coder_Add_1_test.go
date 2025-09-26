package file

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd_1(t *testing.T) {
	type args struct {
		in  int64
		out int64
	}
	tests := []struct {
		name string
		s    *Flow
		args args
		want *Flow
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

			tt.s.Add(tt.args.in, tt.args.out)
			if got := tt.s; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flow.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
