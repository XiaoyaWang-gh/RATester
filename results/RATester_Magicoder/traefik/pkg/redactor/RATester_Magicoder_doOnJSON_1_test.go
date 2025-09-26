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
				input: "https://www.google.com",
			},
			want: "https://www.google.com",
		},
		{
			name: "Test Case 2",
			args: args{
				input: "https://www.facebook.com",
			},
			want: "https://www.facebook.com",
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

			if got := doOnJSON(tt.args.input); got != tt.want {
				t.Errorf("doOnJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
