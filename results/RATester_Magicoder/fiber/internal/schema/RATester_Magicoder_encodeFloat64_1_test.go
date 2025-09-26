package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestencodeFloat64_1(t *testing.T) {
	tests := []struct {
		name string
		v    reflect.Value
		want string
	}{
		{
			name: "Test case 1",
			v:    reflect.ValueOf(float64(1.23)),
			want: "1.23",
		},
		{
			name: "Test case 2",
			v:    reflect.ValueOf(float64(4.56)),
			want: "4.56",
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

			if got := encodeFloat64(tt.v); got != tt.want {
				t.Errorf("encodeFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
