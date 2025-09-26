package hstrings

import (
	"fmt"
	"testing"
)

func TestString_15(t *testing.T) {
	type args struct {
		s StringEqualFold
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{s: "Hello"},
			want: "Hello",
		},
		{
			name: "Test Case 2",
			args: args{s: "WORLD"},
			want: "WORLD",
		},
		{
			name: "Test Case 3",
			args: args{s: "1234567890"},
			want: "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
