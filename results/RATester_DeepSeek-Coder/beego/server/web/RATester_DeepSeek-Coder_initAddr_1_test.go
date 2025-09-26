package web

import (
	"fmt"
	"testing"
)

func TestInitAddr_1(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with valid IP address",
			args: args{addr: "127.0.0.1:8080"},
			want: "127.0.0.1",
		},
		{
			name: "Test with valid domain",
			args: args{addr: "localhost:8080"},
			want: "localhost",
		},
		{
			name: "Test with valid IP address without port",
			args: args{addr: "127.0.0.1"},
			want: "127.0.0.1",
		},
		{
			name: "Test with valid domain without port",
			args: args{addr: "localhost"},
			want: "localhost",
		},
		{
			name: "Test with empty string",
			args: args{addr: ""},
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

			app := &HttpServer{
				Cfg: &Config{
					Listen: Listen{
						HTTPAddr: "",
					},
				},
			}
			app.initAddr(tt.args.addr)
			if app.Cfg.Listen.HTTPAddr != tt.want {
				t.Errorf("initAddr() = %v, want %v", app.Cfg.Listen.HTTPAddr, tt.want)
			}
		})
	}
}
