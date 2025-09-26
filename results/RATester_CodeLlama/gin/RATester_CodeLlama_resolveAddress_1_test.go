package gin

import (
	"fmt"
	"testing"
)

func TestResolveAddress_1(t *testing.T) {
	type args struct {
		addr []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				addr: []string{},
			},
			want: ":8080",
		},
		{
			name: "case 2",
			args: args{
				addr: []string{"127.0.0.1"},
			},
			want: "127.0.0.1",
		},
		{
			name: "case 3",
			args: args{
				addr: []string{"127.0.0.1", "127.0.0.2"},
			},
			want: "127.0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := resolveAddress(tt.args.addr); got != tt.want {
				t.Errorf("resolveAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
