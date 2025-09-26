package npm

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestAdd_16(t *testing.T) {
	type args struct {
		source string
		r      io.Reader
	}
	tests := []struct {
		name string
		b    *packageBuilder
		args args
		want *packageBuilder
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

			if got := tt.b.Add(tt.args.source, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("packageBuilder.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
