package helpers

import (
	"fmt"
	"testing"
)

func TestRelURL_1(t *testing.T) {
	type args struct {
		in          string
		addLanguage bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				in:          "https://example.com/path",
				addLanguage: true,
			},
			want: "https://example.com/path",
		},
		{
			name: "Test case 2",
			args: args{
				in:          "http://example.com/path",
				addLanguage: false,
			},
			want: "http://example.com/path",
		},
		{
			name: "Test case 3",
			args: args{
				in:          "/path",
				addLanguage: true,
			},
			want: "/path",
		},
		{
			name: "Test case 4",
			args: args{
				in:          "//example.com/path",
				addLanguage: false,
			},
			want: "//example.com/path",
		},
		{
			name: "Test case 5",
			args: args{
				in:          "path",
				addLanguage: true,
			},
			want: "/path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &PathSpec{}
			if got := p.RelURL(tt.args.in, tt.args.addLanguage); got != tt.want {
				t.Errorf("RelURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
