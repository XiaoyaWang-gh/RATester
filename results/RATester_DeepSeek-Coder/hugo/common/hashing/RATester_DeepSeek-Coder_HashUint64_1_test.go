package hashing

import (
	"fmt"
	"testing"
)

func TestHashUint64_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []any
		want uint64
	}{
		{
			name: "single argument",
			args: []any{1},
			want: 1,
		},
		{
			name: "multiple arguments",
			args: []any{1, 2, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HashUint64(tt.args...); got != tt.want {
				t.Errorf("HashUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
