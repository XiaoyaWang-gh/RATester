package util

import (
	"fmt"
	"testing"
)

func TestGetAuthKey_1(t *testing.T) {
	type args struct {
		token     string
		timestamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				token:     "test_token",
				timestamp: 1609459200,
			},
			want: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
		{
			name: "Test case 2",
			args: args{
				token:     "another_token",
				timestamp: 1609545600,
			},
			want: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetAuthKey(tt.args.token, tt.args.timestamp); got != tt.want {
				t.Errorf("GetAuthKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
