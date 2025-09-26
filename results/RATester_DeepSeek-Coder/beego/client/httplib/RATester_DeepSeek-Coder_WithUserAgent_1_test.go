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
			name: "TestWithUserAgent",
			args: args{
				userAgent: "TestUserAgent",
			},
			want: WithUserAgent("TestUserAgent"),
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
