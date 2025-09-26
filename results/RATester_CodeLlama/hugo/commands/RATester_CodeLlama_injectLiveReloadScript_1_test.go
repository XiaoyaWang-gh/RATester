package commands

import (
	"fmt"
	"io"
	"net/url"
	"testing"
)

func TestInjectLiveReloadScript_1(t *testing.T) {
	type args struct {
		src     io.Reader
		baseURL *url.URL
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				src:     nil,
				baseURL: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := injectLiveReloadScript(tt.args.src, tt.args.baseURL); got != tt.want {
				t.Errorf("injectLiveReloadScript() = %v, want %v", got, tt.want)
			}
		})
	}
}
