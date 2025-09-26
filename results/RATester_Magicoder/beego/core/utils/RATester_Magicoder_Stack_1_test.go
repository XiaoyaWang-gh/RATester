package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStack_1(t *testing.T) {
	tests := []struct {
		name   string
		skip   int
		indent string
		want   []byte
	}{
		{
			name:   "Test case 1",
			skip:   0,
			indent: "",
			want:   []byte("at TestStack() [stack_test.go:15]\n"),
		},
		{
			name:   "Test case 2",
			skip:   1,
			indent: "",
			want:   []byte("at TestStack() [stack_test.go:15]\n"),
		},
		{
			name:   "Test case 3",
			skip:   0,
			indent: "  ",
			want:   []byte("  at TestStack() [stack_test.go:15]\n"),
		},
		{
			name:   "Test case 4",
			skip:   1,
			indent: "  ",
			want:   []byte("  at TestStack() [stack_test.go:15]\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Stack(tt.skip, tt.indent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack() = %v, want %v", got, tt.want)
			}
		})
	}
}
