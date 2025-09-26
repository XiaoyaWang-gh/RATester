package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithHeader_1(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want BeegoHTTPRequestOption
	}{
		{
			name: "Test WithHeader",
			args: args{
				key:   "Content-Type",
				value: "application/json",
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

			if got := WithHeader(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
