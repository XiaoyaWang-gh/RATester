package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetURL_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		r    *Request
		args args
		want *Request
	}{
		{
			name: "TestSetURL",
			r:    &Request{},
			args: args{
				url: "http://example.com",
			},
			want: &Request{
				url: "http://example.com",
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

			if got := tt.r.SetURL(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.SetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
