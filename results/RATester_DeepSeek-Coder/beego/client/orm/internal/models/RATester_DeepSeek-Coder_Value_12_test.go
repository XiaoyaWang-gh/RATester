package models

import (
	"fmt"
	"testing"
)

func TestValue_12(t *testing.T) {
	tests := []struct {
		name string
		e    BooleanField
		want bool
	}{
		{
			name: "Test case 1",
			e:    BooleanField(true),
			want: true,
		},
		{
			name: "Test case 2",
			e:    BooleanField(false),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.e.Value(); got != tt.want {
				t.Errorf("BooleanField.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
