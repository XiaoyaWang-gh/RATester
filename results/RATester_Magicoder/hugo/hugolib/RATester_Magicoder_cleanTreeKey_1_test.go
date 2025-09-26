package hugolib

import (
	"fmt"
	"testing"
)

func TestcleanTreeKey_1(t *testing.T) {
	tests := []struct {
		name string
		elem []string
		want string
	}{
		{
			name: "Test case 1",
			elem: []string{"test", "case", "1"},
			want: "/test/case/1",
		},
		{
			name: "Test case 2",
			elem: []string{"test", "case", "2"},
			want: "/test/case/2",
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

			if got := cleanTreeKey(tt.elem...); got != tt.want {
				t.Errorf("cleanTreeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
