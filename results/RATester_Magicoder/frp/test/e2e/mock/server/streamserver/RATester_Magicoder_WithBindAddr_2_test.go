package streamserver

import (
	"fmt"
	"testing"
)

func TestWithBindAddr_2(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with valid address",
			args: args{addr: "localhost:8080"},
			want: "localhost:8080",
		},
		{
			name: "Test with empty address",
			args: args{addr: ""},
			want: "",
		},
		{
			name: "Test with invalid address",
			args: args{addr: "localhost:8080:8080"},
			want: "localhost:8080:8080",
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
