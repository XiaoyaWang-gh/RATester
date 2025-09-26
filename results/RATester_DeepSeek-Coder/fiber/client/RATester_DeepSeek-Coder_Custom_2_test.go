package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCustom_2(t *testing.T) {
	type args struct {
		url    string
		method string
	}
	tests := []struct {
		name    string
		r       *Request
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "TestCustom",
			r:    &Request{},
			args: args{
				url:    "http://example.com",
				method: "GET",
			},
			want:    &Response{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.r.Custom(tt.args.url, tt.args.method)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.Custom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.Custom() = %v, want %v", got, tt.want)
			}
		})
	}
}
