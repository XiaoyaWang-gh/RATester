package proxy

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestBuildHttpsRequest_1(t *testing.T) {
	type args struct {
		hostName string
	}
	tests := []struct {
		name string
		args args
		want *http.Request
	}{
		{
			name: "case1",
			args: args{hostName: "www.baidu.com"},
			want: &http.Request{
				RequestURI: "/",
				URL:        &url.URL{Scheme: "https"},
				Host:       "www.baidu.com",
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

			if got := buildHttpsRequest(tt.args.hostName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildHttpsRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
