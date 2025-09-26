package api

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoroutines_1(t *testing.T) {
	tests := []struct {
		name string
		want interface{}
	}{
		{
			name: "Test case 1",
			want: 1,
		},
		{
			name: "Test case 2",
			want: 2,
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

			if got := goroutines(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goroutines() = %v, want %v", got, tt.want)
			}
		})
	}
}
