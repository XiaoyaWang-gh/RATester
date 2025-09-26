package api

import (
	"fmt"
	"testing"
)

func TestGetProviderName_2(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{id: "123@example.com"},
			want: "example.com",
		},
		{
			name: "Test case 2",
			args: args{id: "456@example.net"},
			want: "example.net",
		},
		{
			name: "Test case 3",
			args: args{id: "789@example.org"},
			want: "example.org",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getProviderName(tt.args.id); got != tt.want {
				t.Errorf("getProviderName() = %v, want %v", got, tt.want)
			}
		})
	}
}
