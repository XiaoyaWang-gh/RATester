package tplimpl

import (
	"fmt"
	"testing"
)

func TestRemoveLeadingBOM_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test with leading BOM",
			arg:  "\ufeffHello, World",
			want: "Hello, World",
		},
		{
			name: "Test with no leading BOM",
			arg:  "Hello, World",
			want: "Hello, World",
		},
		{
			name: "Test with empty string",
			arg:  "",
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

			if got := removeLeadingBOM(tt.arg); got != tt.want {
				t.Errorf("removeLeadingBOM() = %v, want %v", got, tt.want)
			}
		})
	}
}
