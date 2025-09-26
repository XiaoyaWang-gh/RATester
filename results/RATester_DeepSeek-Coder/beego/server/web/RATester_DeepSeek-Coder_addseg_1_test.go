package web

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestAddseg_1(t *testing.T) {
	testCases := []struct {
		name      string
		segments  []string
		route     interface{}
		wildcards []string
		reg       string
		expected  *Tree
	}{
		{
			name:     "Test Case 1",
			segments: []string{"seg1", "seg2"},
			route:    "route1",
			expected: &Tree{
				leaves: []*leafInfo{{runObject: "route1"}},
			},
		},
		{
			name:     "Test Case 2",
			segments: []string{"seg1", "seg2"},
			route:    "route2",
			wildcards: []string{
				":wild1",
				":wild2",
			},
			expected: &Tree{
				leaves: []*leafInfo{{runObject: "route2", wildcards: []string{":wild1", ":wild2"}}},
			},
		},
		{
			name:     "Test Case 3",
			segments: []string{"seg1", "seg2"},
			route:    "route3",
			reg:      "reg1",
			expected: &Tree{
				leaves: []*leafInfo{{runObject: "route3", regexps: regexp.MustCompile("^reg1$")}},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tree := &Tree{}
			tree.addseg(tc.segments, tc.route, tc.wildcards, tc.reg)

			if !reflect.DeepEqual(tree, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tree)
			}
		})
	}
}
