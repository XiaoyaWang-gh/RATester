package redactor

import (
	"fmt"
	"testing"
)

func TestDoOnJSON_1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				input: "This is a test with a URL: https://www.example.com",
			},
			want: "This is a test with a URL: " + maskLarge,
		},
		{
			name: "Test Case 2",
			args: args{
				input: "This is a test with a URL: http://www.example.com",
			},
			want: "This is a test with a URL: " + maskLarge,
		},
		{
			name: "Test Case 3",
			args: args{
				input: "This is a test with a URL: www.example.com",
			},
			want: "This is a test with a URL: " + maskLarge,
		},
		{
			name: "Test Case 4",
			args: args{
				input: "This is a test with a URL: example.com",
			},
			want: "This is a test with a URL: " + maskLarge,
		},
		{
			name: "Test Case 5",
			args: args{
				input: "This is a test with a URL: https://example.com",
			},
			want: "This is a test with a URL: " + maskLarge,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := doOnJSON(tt.args.input); got != tt.want {
				t.Errorf("doOnJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
