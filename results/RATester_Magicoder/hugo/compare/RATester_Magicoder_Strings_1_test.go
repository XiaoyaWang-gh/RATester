package compare

import (
	"fmt"
	"testing"
)

func TestStrings_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{
			name: "Test case 1",
			s:    "Hello",
			t:    "hello",
			want: 0,
		},
		{
			name: "Test case 2",
			s:    "World",
			t:    "world",
			want: 0,
		},
		{
			name: "Test case 3",
			s:    "Golang",
			t:    "golang",
			want: 0,
		},
		{
			name: "Test case 4",
			s:    "Test",
			t:    "test",
			want: 0,
		},
		{
			name: "Test case 5",
			s:    "Unit",
			t:    "unit",
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

			if got := Strings(tt.s, tt.t); got != tt.want {
				t.Errorf("Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}
