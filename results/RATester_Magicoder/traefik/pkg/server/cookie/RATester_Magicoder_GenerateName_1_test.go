package cookie

import (
	"fmt"
	"testing"
)

func TestGenerateName_1(t *testing.T) {
	tests := []struct {
		name        string
		backendName string
		want        string
	}{
		{
			name:        "Test case 1",
			backendName: "backend1",
			want:        "_4e1243bd71c5",
		},
		{
			name:        "Test case 2",
			backendName: "backend2",
			want:        "_4e1243bd71c5",
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

			if got := GenerateName(tt.backendName); got != tt.want {
				t.Errorf("GenerateName() = %v, want %v", got, tt.want)
			}
		})
	}
}
