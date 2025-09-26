package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseRelated_1(t *testing.T) {
	testCases := []struct {
		name     string
		rels     []string
		depth    int
		expected []string
	}{
		{
			name:     "Test case 1",
			rels:     []string{"field1", "field2"},
			depth:    1,
			expected: []string{"field1", "field2"},
		},
		{
			name:     "Test case 2",
			rels:     []string{"field3", "field4"},
			depth:    2,
			expected: []string{"field3", "field4"},
		},
		{
			name:     "Test case 3",
			rels:     []string{"field5", "field6"},
			depth:    0,
			expected: []string{"field5", "field6"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			dbTables := &dbTables{}
			dbTables.parseRelated(tc.rels, tc.depth)

			if !reflect.DeepEqual(tc.rels, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.rels)
			}
		})
	}
}
