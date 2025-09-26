package hstrings

import (
	"fmt"
	"testing"
)

func TestEqualAny_1(t *testing.T) {
	type args struct {
		a string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				a: "test",
				b: []string{"test", "test2", "test3"},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				a: "test4",
				b: []string{"test", "test2", "test3"},
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

			if got := EqualAny(tt.args.a, tt.args.b...); got != tt.want {
				t.Errorf("EqualAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
