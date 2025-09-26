package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_6(t *testing.T) {
	min := Min{Min: 10, Key: "test"}

	tests := []struct {
		name string
		obj  interface{}
		want bool
	}{
		{"int64", int64(15), true},
		{"int", 15, true},
		{"int32", int32(15), true},
		{"int16", int16(15), true},
		{"int8", int8(15), true},
		{"int64_less", int64(5), false},
		{"int_less", 5, false},
		{"int32_less", int32(5), false},
		{"int16_less", int16(5), false},
		{"int8_less", int8(5), false},
		{"string", "test", false},
		{"float64", float64(15), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := min.IsSatisfied(tt.obj); got != tt.want {
				t.Errorf("IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
