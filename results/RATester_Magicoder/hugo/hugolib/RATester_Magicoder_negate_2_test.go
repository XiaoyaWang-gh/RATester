package hugolib

import (
	"fmt"
	"testing"

	"github.com/frankban/quicktest"
)

func Testnegate_2(t *testing.T) {
	b := &IntegrationTestBuilder{
		C: &quicktest.C{},
	}

	tests := []struct {
		name     string
		match    string
		expected string
		negate   bool
	}{
		{
			name:     "positive case",
			match:    "! hello",
			expected: "hello",
			negate:   true,
		},
		{
			name:     "negative case",
			match:    "hello",
			expected: "hello",
			negate:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			match, negate := b.negate(test.match)
			if match != test.expected {
				t.Errorf("Expected match to be %s, but got %s", test.expected, match)
			}
			if negate != test.negate {
				t.Errorf("Expected negate to be %t, but got %t", test.negate, negate)
			}
		})
	}
}
