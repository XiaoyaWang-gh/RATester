package doctree

import (
	"fmt"
	"testing"
)

func TestcleanKey_2(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want string
	}{
		{
			name: "Test case 1",
			key:  "/",
			want: "",
		},
		{
			name: "Test case 2",
			key:  "test",
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cleanKey(tt.key); got != tt.want {
				t.Errorf("cleanKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
