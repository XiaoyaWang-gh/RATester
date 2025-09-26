package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetBaseURL_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		c    *Client
		args args
		want *Client
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

			if got := tt.c.SetBaseURL(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SetBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
