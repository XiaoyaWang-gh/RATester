package cookie

import (
	"fmt"
	"testing"
)

func TestGetName_4(t *testing.T) {
	tests := []struct {
		name        string
		cookieName  string
		backendName string
		want        string
	}{
		{
			name:        "Test case 1",
			cookieName:  "test_cookie",
			backendName: "test_backend",
			want:        "test_cookie",
		},
		{
			name:        "Test case 2",
			cookieName:  "",
			backendName: "test_backend",
			want:        "test_backend",
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

			if got := GetName(tt.cookieName, tt.backendName); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
