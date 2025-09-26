package mock

import (
	"fmt"
	"io/ioutil"
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
			name: "Test case 1",
			args: args{
				path: "/test",
				resp: &http.Response{
					Status:     "200 OK",
					StatusCode: 200,
					Proto:      "HTTP/1.0",
					ProtoMajor: 1,
					ProtoMinor: 0,
					Header:     http.Header{},
					Body:       ioutil.NopCloser(strings.NewReader("test")),
					Request:    &http.Request{},
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
					Header:     http.Header{},
					Body:       ioutil.NopCloser(strings.NewReader("test")),
					Request:    &http.Request{},
				},
				err: nil,
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewMockByPath(tt.args.path, tt.args.resp, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockByPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
