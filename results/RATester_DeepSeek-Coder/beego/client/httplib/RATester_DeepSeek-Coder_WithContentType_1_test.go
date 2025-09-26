package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithContentType_1(t *testing.T) {
	type args struct {
		contentType string
	}
	tests := []struct {
		name string
		args args
		want BeegoHTTPRequestOption
	}{
		{
			name: "TestWithContentType",
			args: args{
				contentType: "application/json",
			},
			want: func(request *BeegoHTTPRequest) {
				request.Header("Content-Type", "application/json")
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

			if got := WithContentType(tt.args.contentType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}
