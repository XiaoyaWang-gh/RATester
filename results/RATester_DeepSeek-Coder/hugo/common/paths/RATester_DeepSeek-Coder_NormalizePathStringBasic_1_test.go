package paths

import (
	"fmt"
	"testing"
)

func TestNormalizePathStringBasic_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with a string with spaces",
			args: args{s: "This is a test string"},
			want: "this-is-a-test-string",
		},
		{
			name: "Test with a string with special characters",
			args: args{s: "This*is@a#test$string%^&*()_+"},
			want: "this-is-a-test-string",
		},
		{
			name: "Test with a string with numbers",
			args: args{s: "This1is2a3test4string"},
			want: "this1is2a3test4string",
		},
		{
			name: "Test with a string with uppercase letters",
			args: args{s: "ThisIsATestString"},
			want: "thisisateststring",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NormalizePathStringBasic(tt.args.s); got != tt.want {
				t.Errorf("NormalizePathStringBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}
