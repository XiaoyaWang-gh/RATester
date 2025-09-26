package kinds

import (
	"fmt"
	"testing"
)

func TestGetKindAny_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test case 1",
			arg:  "field name",
			want: "field want",
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

			if got := GetKindAny(tt.arg); got != tt.want {
				t.Errorf("GetKindAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
