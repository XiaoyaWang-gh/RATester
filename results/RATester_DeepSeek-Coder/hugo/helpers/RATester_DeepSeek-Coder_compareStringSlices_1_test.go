package helpers

import (
	"fmt"
	"testing"
)

func TestCompareStringSlices_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		a, b []string
		want bool
	}

	tests := []test{
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b"}, false},
		{nil, nil, true},
		{[]string{}, []string{}, true},
		{nil, []string{}, false},
		{[]string{}, nil, false},
	}

	for _, tt := range tests {
		got := compareStringSlices(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("compareStringSlices(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
