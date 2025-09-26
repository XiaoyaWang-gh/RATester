package codegen

import (
	"fmt"
	"testing"
)

func TestFirstToUpper_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with lowercase string",
			args: args{name: "test"},
			want: "Test",
		},
		{
			name: "Test with uppercase string",
			args: args{name: "TEST"},
			want: "TEST",
		},
		{
			name: "Test with mixed case string",
			args: args{name: "tEsT"},
			want: "TEST",
		},
		{
			name: "Test with number",
			args: args{name: "123test"},
			want: "123test",
		},
		{
			name: "Test with special character",
			args: args{name: "@test"},
			want: "@test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := firstToUpper(tt.args.name); got != tt.want {
				t.Errorf("firstToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
