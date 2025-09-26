package wrr

import (
	"fmt"
	"testing"
)

func TestLess_2(t *testing.T) {
	b := &Balancer{
		handlers: []*namedHandler{
			{name: "handler1", deadline: 1.0},
			{name: "handler2", deadline: 2.0},
			{name: "handler3", deadline: 3.0},
		},
	}

	tests := []struct {
		name     string
		i        int
		j        int
		expected bool
	}{
		{"less than", 0, 1, true},
		{"greater than", 1, 0, false},
		{"equal", 0, 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := b.Less(test.i, test.j)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
