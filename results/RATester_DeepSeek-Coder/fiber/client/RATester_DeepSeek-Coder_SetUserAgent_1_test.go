package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetUserAgent_1(t *testing.T) {
	type args struct {
		ua string
	}
	tests := []struct {
		name string
		r    *Request
		args args
		want *Request
	}{
		{
			name: "TestSetUserAgent",
			r:    &Request{},
			args: args{ua: "TestUserAgent"},
			want: &Request{userAgent: "TestUserAgent"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.r.SetUserAgent(tt.args.ua); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.SetUserAgent() = %v, want %v", got, tt.want)
			}
		})
	}
}
