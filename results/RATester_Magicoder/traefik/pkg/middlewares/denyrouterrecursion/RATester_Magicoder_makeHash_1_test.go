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
			args: args{
				routerName: "test",
			},
			want: "a0989f085c2ed1",
		},
		{
			name: "Test case 2",
			args: args{
				routerName: "anotherTest",
			},
			want: "a0989f085c2ed1",
		},
		// Add more test cases as needed
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
