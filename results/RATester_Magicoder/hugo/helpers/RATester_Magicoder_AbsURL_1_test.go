package helpers

import (
	"fmt"
	"testing"
)

func TestAbsURL_1(t *testing.T) {
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
				in:          "test.html",
				addLanguage: true,
			},
			want: "test.html",
		},
		{
			name: "Test case 2",
			args: args{
				in:          "http://example.com",
				addLanguage: false,
			},
			want: "http://example.com",
		},
		{
			name: "Test case 3",
			args: args{
				in:          "//example.com",
				addLanguage: true,
			},
			want: "//example.com",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &PathSpec{}
			if got := p.AbsURL(tt.args.in, tt.args.addLanguage); got != tt.want {
				t.Errorf("AbsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
