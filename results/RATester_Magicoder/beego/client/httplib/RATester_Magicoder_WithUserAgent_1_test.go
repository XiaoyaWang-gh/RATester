package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithUserAgent_1(t *testing.T) {
	type args struct {
		userAgent string
	}
	tests := []struct {
		name string
		args args
		want ClientOption
	}{
		{
			name: "Test WithUserAgent",
			args: args{
				userAgent: "test-user-agent",
			},
			want: func(client *Client) {
				client.Setting.UserAgent = "test-user-agent"
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

			if got := WithUserAgent(tt.args.userAgent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithUserAgent() = %v, want %v", got, tt.want)
			}
		})
	}
}
