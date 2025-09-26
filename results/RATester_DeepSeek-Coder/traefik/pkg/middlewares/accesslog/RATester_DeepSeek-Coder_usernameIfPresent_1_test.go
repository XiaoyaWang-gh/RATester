package accesslog

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUsernameIfPresent_1(t *testing.T) {
	type args struct {
		theURL *url.URL
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with nil user",
			args: args{
				theURL: &url.URL{
					User: nil,
				},
			},
			want: "-",
		},
		{
			name: "Test with empty username",
			args: args{
				theURL: &url.URL{
					User: &url.Userinfo{},
				},
			},
			want: "-",
		},
		{
			name: "Test with non-empty username",
			args: args{
				theURL: &url.URL{
					User: url.UserPassword("username", "password"),
				},
			},
			want: "username",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := usernameIfPresent(tt.args.theURL); got != tt.want {
				t.Errorf("usernameIfPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}
