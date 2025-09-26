package compare

import (
	"fmt"
	"testing"
)

func TestcompareGet_1(t *testing.T) {
	ns := &Namespace{}
	tests := []struct {
		name string
		a    any
		b    any
		want float64
	}{
		{
			name: "Test case 1",
			a:    1,
			b:    2,
			want: 1,
		},
		{
			name: "Test case 2",
			a:    3,
			b:    1,
			want: 1,
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

			got, _ := ns.compareGet(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("compareGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
