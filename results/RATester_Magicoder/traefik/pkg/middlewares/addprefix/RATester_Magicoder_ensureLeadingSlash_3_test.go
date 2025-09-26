package addprefix

import (
	"fmt"
	"testing"
)

func TestEnsureLeadingSlash_3(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Empty string",
			str:  "",
			want: "",
		},
		{
			name: "String with leading slash",
			str:  "/test",
			want: "/test",
		},
		{
			name: "String without leading slash",
			str:  "test",
			want: "/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ensureLeadingSlash(tt.str); got != tt.want {
				t.Errorf("ensureLeadingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
