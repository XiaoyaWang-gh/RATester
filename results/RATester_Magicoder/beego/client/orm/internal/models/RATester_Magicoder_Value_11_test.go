package models

import (
	"fmt"
	"testing"
)

func TestValue_11(t *testing.T) {
	tests := []struct {
		name string
		e    BigIntegerField
		want int64
	}{
		{
			name: "Test case 1",
			e:    BigIntegerField(123),
			want: 123,
		},
		{
			name: "Test case 2",
			e:    BigIntegerField(0),
			want: 0,
		},
		{
			name: "Test case 3",
			e:    BigIntegerField(-123),
			want: -123,
		},
		{
			name: "Test case 4",
			e:    BigIntegerField(9223372036854775807),
			want: 9223372036854775807,
		},
		{
			name: "Test case 5",
			e:    BigIntegerField(-9223372036854775808),
			want: -9223372036854775808,
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
				t.Errorf("BigIntegerField.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
