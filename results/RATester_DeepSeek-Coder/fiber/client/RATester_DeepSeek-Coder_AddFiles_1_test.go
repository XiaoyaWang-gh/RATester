package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddFiles_1(t *testing.T) {
	type args struct {
		files []*File
	}
	tests := []struct {
		name string
		r    *Request
		args args
		want *Request
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

			if got := tt.r.AddFiles(tt.args.files...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.AddFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
