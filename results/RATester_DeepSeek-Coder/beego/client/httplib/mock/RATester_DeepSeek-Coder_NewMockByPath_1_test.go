package mock

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestNewMockByPath_1(t *testing.T) {
	type args struct {
		path string
		resp *http.Response
		err  error
	}
	tests := []struct {
		name string
		args args
		want *Mock
	}{
		{
			name: "TestNewMockByPath",
			args: args{
				path: "/test",
				resp: &http.Response{
					Status:     "200 OK",
					StatusCode: 200,
					Proto:      "HTTP/1.0",
					ProtoMajor: 1,
					ProtoMinor: 0,
					Header:     make(http.Header),
					Body:       io.NopCloser(strings.NewReader("")),
				},
				err: nil,
			},
			want: &Mock{
				cond: NewSimpleCondition("/test"),
				resp: &http.Response{
					Status:     "200 OK",
					StatusCode: 200,
					Proto:      "HTTP/1.0",
					ProtoMajor: 1,
					ProtoMinor: 0,
					Header:     make(http.Header),
					Body:       io.NopCloser(strings.NewReader("")),
				},
				err: nil,
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

			got := NewMockByPath(tt.args.path, tt.args.resp, tt.args.err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockByPath() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.cond, tt.want.cond) {
				t.Errorf("NewMockByPath().cond = %v, want %v", got.cond, tt.want.cond)
			}
		})
	}
}
