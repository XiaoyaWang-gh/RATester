package httplib

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestNewHttpResponseWithJsonBody_1(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want *http.Response
	}{
		{
			name: "Test with string",
			args: args{data: "test"},
			want: &http.Response{
				ContentLength: 4,
				Body:          io.NopCloser(bytes.NewReader([]byte("test"))),
			},
		},
		{
			name: "Test with []byte",
			args: args{data: []byte("test")},
			want: &http.Response{
				ContentLength: 4,
				Body:          io.NopCloser(bytes.NewReader([]byte("test"))),
			},
		},
		{
			name: "Test with struct",
			args: args{data: struct{ Name string }{Name: "test"}},
			want: &http.Response{
				ContentLength: 15,
				Body:          io.NopCloser(bytes.NewReader([]byte(`{"Name":"test"}`))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewHttpResponseWithJsonBody(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpResponseWithJsonBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
