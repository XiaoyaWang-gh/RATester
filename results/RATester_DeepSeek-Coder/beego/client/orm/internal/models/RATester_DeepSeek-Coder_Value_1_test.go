package models

import (
	"fmt"
	"testing"
)

func TestValue_1(t *testing.T) {
	tests := []struct {
		name string
		e    FloatField
		want float64
	}{
		{
			name: "Test case 1",
			e:    FloatField(123.456),
			want: 123.456,
		},
		{
			name: "Test case 2",
			e:    FloatField(-123.456),
			want: -123.456,
		},
		{
			name: "Test case 3",
			e:    FloatField(0),
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
				t.Errorf("FloatField.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
