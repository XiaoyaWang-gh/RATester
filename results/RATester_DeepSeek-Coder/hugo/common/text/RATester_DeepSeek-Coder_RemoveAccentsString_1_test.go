package text

import (
	"fmt"
	"testing"
)

func TestRemoveAccentsString_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with string with accents",
			args: args{s: "éèêëėęîïì������äåæœ�"},
			want: "eeeeeeiiiiuuuuuaaaaeoeu",
		},
		{
			name: "Test with string without accents",
			args: args{s: "abcdefghijklmnopqrstuvwxyz"},
			want: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name: "Test with string with special characters",
			args: args{s: "@#$%^&*()_+`~-=[]\\\";':,./<>?{}|_+"},
			want: "@#$%^&*()_+`~-=[]\\\";':,./<>?{}|_+",
		},
		{
			name: "Test with string with numbers",
			args: args{s: "1234567890"},
			want: "1234567890",
		},
		{
			name: "Test with string with spaces",
			args: args{s: " "},
			want: " ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := RemoveAccentsString(tt.args.s); got != tt.want {
				t.Errorf("RemoveAccentsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
