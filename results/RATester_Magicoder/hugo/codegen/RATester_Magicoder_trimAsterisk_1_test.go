package codegen

import (
	"fmt"
	"testing"
)

func TesttrimAsterisk_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test trimAsterisk with asterisk",
			args: args{name: "*test"},
			want: "test",
		},
		{
			name: "Test trimAsterisk without asterisk",
			args: args{name: "test"},
			want: "test",
		},
		{
			name: "Test trimAsterisk with empty string",
			args: args{name: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := trimAsterisk(tt.args.name); got != tt.want {
				t.Errorf("trimAsterisk() = %v, want %v", got, tt.want)
			}
		})
	}
}
