package compare

import (
	"fmt"
	"testing"
)

func TestcompareTwoUints_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name string
		a    uint64
		b    uint64
		want float64
	}{
		{
			name: "a < b",
			a:    1,
			b:    2,
			want: 1,
		},
		{
			name: "a == b",
			a:    2,
			b:    2,
			want: 0,
		},
		{
			name: "a > b",
			a:    3,
			b:    2,
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

			got, _ := ns.compareTwoUints(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("compareTwoUints() = %v, want %v", got, tt.want)
			}
		})
	}
}
