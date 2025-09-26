package wrr

import (
	"fmt"
	"testing"
)

func TestHash_1(t *testing.T) {
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
			args: args{input: "test"},
			want: "a0982d2222c3a082a900000000000000",
		},
		{
			name: "Test Case 2",
			args: args{input: "hello"},
			want: "30982d2222c3a082a900000000000000",
		},
		{
			name: "Test Case 3",
			args: args{input: "world"},
			want: "40982d2222c3a082a900000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := hash(tt.args.input); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
