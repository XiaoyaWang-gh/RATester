package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExpectOneOf_1(t *testing.T) {
	tree := &Tree{
		Name: "test",
	}

	expected1 := item{
		typ:  itemType(1),
		pos:  1,
		val:  "expected1",
		line: 1,
	}

	expected2 := item{
		typ:  itemType(2),
		pos:  2,
		val:  "expected2",
		line: 2,
	}

	tests := []struct {
		name      string
		expected1 itemType
		expected2 itemType
		context   string
		want      item
	}{
		{
			name:      "test1",
			expected1: itemType(1),
			expected2: itemType(2),
			context:   "test context",
			want:      expected1,
		},
		{
			name:      "test2",
			expected1: itemType(1),
			expected2: itemType(2),
			context:   "test context",
			want:      expected2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tree.expectOneOf(tt.expected1, tt.expected2, tt.context)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expectOneOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
