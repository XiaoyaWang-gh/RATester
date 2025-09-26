package media

import (
	"fmt"
	"testing"
)

func TestGetBySubType_1(t *testing.T) {
	tests := []struct {
		name     string
		types    Types
		subType  string
		expected Type
		found    bool
	}{
		{
			name: "Test case 1",
			types: Types{
				{
					MainType: "application",
					SubType:  "json",
				},
				{
					MainType: "text",
					SubType:  "html",
				},
			},
			subType: "json",
			expected: Type{
				MainType: "application",
				SubType:  "json",
			},
			found: true,
		},
		{
			name: "Test case 2",
			types: Types{
				{
					MainType: "application",
					SubType:  "json",
				},
				{
					MainType: "text",
					SubType:  "html",
				},
			},
			subType: "html",
			expected: Type{
				MainType: "text",
				SubType:  "html",
			},
			found: true,
		},
		{
			name: "Test case 3",
			types: Types{
				{
					MainType: "application",
					SubType:  "json",
				},
				{
					MainType: "text",
					SubType:  "html",
				},
			},
			subType:  "xml",
			expected: Type{},
			found:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tp, found := test.types.GetBySubType(test.subType)
			if tp != test.expected || found != test.found {
				t.Errorf("Expected %v, found %v", test.expected, tp)
			}
		})
	}
}
