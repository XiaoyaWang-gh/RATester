package models

import (
	"fmt"
	"testing"
)

func TestValue_6(t *testing.T) {
	tests := []struct {
		name string
		e    PositiveSmallIntegerField
		want uint16
	}{
		{
			name: "Test case 1",
			e:    10,
			want: 10,
		},
		{
			name: "Test case 2",
			e:    20,
			want: 20,
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

			if got := tt.e.Value(); got != tt.want {
				t.Errorf("PositiveSmallIntegerField.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
