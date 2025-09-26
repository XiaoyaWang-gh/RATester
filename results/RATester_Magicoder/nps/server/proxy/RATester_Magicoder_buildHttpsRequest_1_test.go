package proxy

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestbuildHttpsRequest_1(t *testing.T) {
	type args struct {
		hostName string
	}
	tests := []struct {
		name string
		args args
		want *http.Request
	}{
		{
			name: "Test Case 1",
			args: args{
				hostName: "example.com",
			},
			want: &http.Request{
				RequestURI: "/",
				URL: &url.URL{
					Scheme: "https",
					Host:   "example.com",
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				hostName: "test.com",
			},
			want: &http.Request{
				RequestURI: "/",
				URL: &url.URL{
					Scheme: "https",
					Host:   "test.com",
				},
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
