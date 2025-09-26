package identity

import (
	"fmt"
	"testing"
)

func TestCleanStringIdentity_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want StringIdentity
	}{
		{
			name: "CleanStringIdentity_EmptyString",
			args: "",
			want: "",
		},
		{
			name: "CleanStringIdentity_SingleSpace",
			args: " ",
			want: "",
		},
		{
			name: "CleanStringIdentity_MultipleSpaces",
			args: "   ",
			want: "",
		},
		{
			name: "CleanStringIdentity_SingleWord",
			args: "Hello",
			want: "Hello",
		},
		{
			name: "CleanStringIdentity_MultipleWords",
			args: "Hello World",
			want: "Hello World",
		},
		{
			name: "CleanStringIdentity_LeadingAndTrailingSpaces",
			args: "  Hello World  ",
			want: "Hello World",
		},
		{
			name: "CleanStringIdentity_SpecialCharacters",
			args: "Hello World!",
			want: "Hello World!",
		},
		{
			name: "CleanStringIdentity_Numbers",
			args: "1234567890",
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

			if got := CleanStringIdentity(tt.args); got != tt.want {
				t.Errorf("CleanStringIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
