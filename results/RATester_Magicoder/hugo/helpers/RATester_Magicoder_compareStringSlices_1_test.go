package helpers

import (
	"fmt"
	"testing"
)

func TestcompareStringSlices_1(t *testing.T) {
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
		{[]string{"a", "b", "c"}, []string{"a", "b"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, false},
		{[]string{"a", "b", "c"}, nil, false},
		{nil, []string{"a", "b", "c"}, false},
		{nil, nil, true},
	}

	for _, tt := range tests {
		got := compareStringSlices(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("compareStringSlices(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
