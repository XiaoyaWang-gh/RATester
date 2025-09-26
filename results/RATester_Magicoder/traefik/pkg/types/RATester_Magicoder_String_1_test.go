package types

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	tests := []struct {
		name string
		f    FileOrContent
		want string
	}{
		{
			name: "Test case 1",
			f:    "test content",
			want: "test content",
		},
		{
			name: "Test case 2",
			f:    "test path",
			want: "test path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.f.String(); got != tt.want {
				t.Errorf("FileOrContent.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
