package common

import (
	"fmt"
	"testing"
)

func TestGetHostByName_1(t *testing.T) {
	type args struct {
		hostname string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with valid domain",
			args: args{hostname: "google.com"},
			want: "216.58.216.46",
		},
		{
			name: "Test with invalid domain",
			args: args{hostname: "invalid.domain"},
			want: "invalid.domain",
		},
		{
			name: "Test with empty string",
			args: args{hostname: ""},
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

			if got := GetHostByName(tt.args.hostname); got != tt.want {
				t.Errorf("GetHostByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
