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
			name: "case1",
			args: args{
				data: "test",
			},
			want: &http.Response{
				ContentLength: int64(len("test")),
				Body:          io.NopCloser(bytes.NewReader([]byte("test"))),
			},
		},
		{
			name: "case2",
			args: args{
				data: []byte("test"),
			},
			want: &http.Response{
				ContentLength: int64(len("test")),
				Body:          io.NopCloser(bytes.NewReader([]byte("test"))),
			},
		},
		{
			name: "case3",
			args: args{
				data: map[string]interface{}{
					"test": "test",
				},
			},
			want: &http.Response{
				ContentLength: int64(len("{\"test\":\"test\"}")),
				Body:          io.NopCloser(bytes.NewReader([]byte("{\"test\":\"test\"}"))),
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
