package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetReferer_1(t *testing.T) {
	type args struct {
		referer string
	}
	tests := []struct {
		name string
		r    *Request
		args args
		want *Request
	}{
		{
			name: "TestSetReferer",
			r:    &Request{},
			args: args{
				referer: "http://example.com",
			},
			want: &Request{
				referer: "http://example.com",
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

			if got := tt.r.SetReferer(tt.args.referer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.SetReferer() = %v, want %v", got, tt.want)
			}
		})
	}
}
