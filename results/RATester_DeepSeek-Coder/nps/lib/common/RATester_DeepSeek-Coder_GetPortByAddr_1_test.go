package common

import (
	"fmt"
	"testing"
)

func TestGetPortByAddr_1(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{addr: "localhost:8080"},
			want: 8080,
		},
		{
			name: "Test case 2",
			args: args{addr: "127.0.0.1:8080"},
			want: 8080,
		},
		{
			name: "Test case 3",
			args: args{addr: "127.0.0.1"},
			want: 0,
		},
		{
			name: "Test case 4",
			args: args{addr: "localhost"},
			want: 0,
		},
		{
			name: "Test case 5",
			args: args{addr: ":8080"},
			want: 8080,
		},
		{
			name: "Test case 6",
			args: args{addr: "8080"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetPortByAddr(tt.args.addr); got != tt.want {
				t.Errorf("GetPortByAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
