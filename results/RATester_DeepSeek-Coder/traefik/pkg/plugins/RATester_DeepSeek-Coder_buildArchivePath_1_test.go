package plugins

import (
	"fmt"
	"testing"
)

func TestBuildArchivePath_1(t *testing.T) {
	type args struct {
		pName    string
		pVersion string
	}
	tests := []struct {
		name string
		c    *Client
		args args
		want string
	}{
		{
			name: "Test case 1",
			c:    &Client{archives: "/tmp/archives"},
			args: args{pName: "test", pVersion: "1.0.0"},
			want: "/tmp/archives/test/1.0.0.zip",
		},
		{
			name: "Test case 2",
			c:    &Client{archives: "/tmp/archives"},
			args: args{pName: "test2", pVersion: "2.0.0"},
			want: "/tmp/archives/test2/2.0.0.zip",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.buildArchivePath(tt.args.pName, tt.args.pVersion); got != tt.want {
				t.Errorf("Client.buildArchivePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
