package requestdecorator

import (
	"fmt"
	"testing"
)

func TestLen_3(t *testing.T) {
	testCases := []struct {
		name string
		a    byTTL
		want int
	}{
		{
			name: "Empty slice",
			a:    byTTL{},
			want: 0,
		},
		{
			name: "Non-empty slice",
			a:    byTTL{&cnameResolv{}, &cnameResolv{}},
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tc.a.Len(); got != tc.want {
				t.Errorf("Len() = %v, want %v", got, tc.want)
			}
		})
	}
}
