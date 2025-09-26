package proxy

import (
	"fmt"
	"net/http"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := buildHttpsRequest(tt.args.hostName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildHttpsRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
