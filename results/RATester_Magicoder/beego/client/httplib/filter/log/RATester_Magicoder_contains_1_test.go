package log

import (
	"fmt"
	"testing"
)

func Testcontains_1(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				s: []string{"apple", "banana", "cherry"},
				e: "banana",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				s: []string{"apple", "banana", "cherry"},
				e: "grape",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				s: []string{},
				e: "grape",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
