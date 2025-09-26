package models

import (
	"fmt"
	"testing"
)

func TestValue_5(t *testing.T) {
	tests := []struct {
		name string
		e    PositiveBigIntegerField
		want uint64
	}{
		{
			name: "Test case 1",
			e:    100,
			want: 100,
		},
		{
			name: "Test case 2",
			e:    18446744073709551615,
			want: 18446744073709551615,
		},
		{
			name: "Test case 3",
			e:    0,
			want: 0,
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
				t.Errorf("PositiveBigIntegerField.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
