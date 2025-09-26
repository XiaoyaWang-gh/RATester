package httpserver

import (
	"fmt"
	"testing"
)

func TestWithBindAddr_1(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestWithBindAddr_1",
			args: args{addr: "127.0.0.1:8080"},
			want: "127.0.0.1:8080",
		},
		{
			name: "TestWithBindAddr_2",
			args: args{addr: "0.0.0.0:8080"},
			want: "0.0.0.0:8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &Server{}
			got := WithBindAddr(tt.args.addr)(s)
			if got.bindAddr != tt.want {
				t.Errorf("WithBindAddr() = %v, want %v", got.bindAddr, tt.want)
			}
		})
	}
}
