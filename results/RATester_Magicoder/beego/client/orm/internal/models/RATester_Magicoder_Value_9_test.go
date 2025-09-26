package models

import (
	"fmt"
	"testing"
)

func TestValue_9(t *testing.T) {
	tests := []struct {
		name string
		e    IntegerField
		want int32
	}{
		{
			name: "Test case 1",
			e:    IntegerField(1),
			want: 1,
		},
		{
			name: "Test case 2",
			e:    IntegerField(0),
			want: 0,
		},
		{
			name: "Test case 3",
			e:    IntegerField(-1),
			want: -1,
		},
		{
			name: "Test case 4",
			e:    IntegerField(2147483647),
			want: 2147483647,
		},
		{
			name: "Test case 5",
			e:    IntegerField(-2147483648),
			want: -2147483648,
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
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
