package tplimpl

import (
	"fmt"
	"testing"
)

func TestremoveLeadingBOM_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Test with leading BOM",
			s:    "\ufeffHello",
			want: "Hello",
		},
		{
			name: "Test without leading BOM",
			s:    "Hello",
			want: "Hello",
		},
		{
			name: "Test with empty string",
			s:    "",
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

			if got := removeLeadingBOM(tt.s); got != tt.want {
				t.Errorf("removeLeadingBOM() = %v, want %v", got, tt.want)
			}
		})
	}
}
