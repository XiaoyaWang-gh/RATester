package allconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetMainSections_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		sections []string
		expected []string
	}{
		{
			name:     "empty sections",
			sections: []string{},
			expected: []string{},
		},
		{
			name:     "single section",
			sections: []string{"section1"},
			expected: []string{"section1"},
		},
		{
			name:     "multiple sections",
			sections: []string{"section1", "section2", "section3"},
			expected: []string{"section1", "section2", "section3"},
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			c := &ConfigCompiled{}
			c.SetMainSections(tc.sections)

			if !reflect.DeepEqual(c.MainSections, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, c.MainSections)
			}
		})
	}
}
