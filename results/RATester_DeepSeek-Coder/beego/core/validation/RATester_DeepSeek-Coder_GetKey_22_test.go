package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_22(t *testing.T) {
	tests := []struct {
		name string
		a    AlphaNumeric
		want string
	}{
		{
			name: "Test case 1",
			a:    AlphaNumeric{Key: "TestKey1"},
			want: "TestKey1",
		},
		{
			name: "Test case 2",
			a:    AlphaNumeric{Key: "TestKey2"},
			want: "TestKey2",
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

			if got := tt.a.GetKey(); got != tt.want {
				t.Errorf("AlphaNumeric.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
