package denyrouterrecursion

import (
	"fmt"
	"testing"
)

func TestMakeHash_1(t *testing.T) {
	type args struct {
		routerName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{routerName: "router1"},
			want: "b324f8f21e5c8f87",
		},
		{
			name: "Test case 2",
			args: args{routerName: "router2"},
			want: "b324f8f21e5c8f88",
		},
		{
			name: "Test case 3",
			args: args{routerName: "router3"},
			want: "b324f8f21e5c8f89",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := makeHash(tt.args.routerName); got != tt.want {
				t.Errorf("makeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
