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
			name: "test buildArchivePath",
			c: &Client{
				archives: "test",
			},
			args: args{
				pName:    "test",
				pVersion: "test",
			},
			want: "test/test/test.zip",
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
				t.Errorf("buildArchivePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
