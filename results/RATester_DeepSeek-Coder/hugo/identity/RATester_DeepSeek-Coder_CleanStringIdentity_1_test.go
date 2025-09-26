package identity

import (
	"fmt"
	"testing"
)

func TestCleanStringIdentity_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want StringIdentity
	}{
		{
			name: "Test case 1",
			arg:  "test",
			want: "test",
		},
		{
			name: "Test case 2",
			arg:  "test2",
			want: "test2",
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

			if got := CleanStringIdentity(tt.arg); got != tt.want {
				t.Errorf("CleanStringIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
