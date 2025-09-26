package compare

import (
	"fmt"
	"testing"
)

func TestStrings_1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "Hello",
				t: "hello",
			},
			want: 0,
		},
		{
			name: "Test Case 2",
			args: args{
				s: "World",
				t: "hello",
			},
			want: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				s: "hello",
				t: "World",
			},
			want: -1,
		},
		{
			name: "Test Case 4",
			args: args{
				s: "Hello",
				t: "Hello",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Strings(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}
