package page

import (
	"fmt"
	"testing"
)

func TestisGlobWithExtension_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{s: "test.txt"},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{s: "test/test.txt"},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{s: "test.txt/test"},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{s: "test"},
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

			if got := isGlobWithExtension(tt.args.s); got != tt.want {
				t.Errorf("isGlobWithExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
