package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestName_10(t *testing.T) {
	tests := []struct {
		name string
		o    OutputFormat
		want string
	}{
		{
			name: "Test Case 1",
			o:    OutputFormat{Format: output.Format{Name: "TestName"}},
			want: "TestName",
		},
		{
			name: "Test Case 2",
			o:    OutputFormat{Format: output.Format{Name: "AnotherName"}},
			want: "AnotherName",
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

			if got := tt.o.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
