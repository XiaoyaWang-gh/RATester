package hugo

import (
	"fmt"
	"testing"
)

func TestString_7(t *testing.T) {
	tests := []struct {
		name string
		h    VersionString
		want string
	}{
		{
			name: "Test case 1",
			h:    "1.0.0",
			want: "1.0.0",
		},
		{
			name: "Test case 2",
			h:    "2.0.0",
			want: "2.0.0",
		},
		{
			name: "Test case 3",
			h:    "3.0.0",
			want: "3.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.h.String(); got != tt.want {
				t.Errorf("VersionString.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
