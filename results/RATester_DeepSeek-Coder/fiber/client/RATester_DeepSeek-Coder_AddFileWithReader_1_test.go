package client

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestAddFileWithReader_1(t *testing.T) {
	type args struct {
		name   string
		reader io.ReadCloser
	}
	tests := []struct {
		name string
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

			r := &Request{}
			if got := r.AddFileWithReader(tt.args.name, tt.args.reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.AddFileWithReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
