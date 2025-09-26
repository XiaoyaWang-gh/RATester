package metadecoders

import (
	"fmt"
	"testing"
)

func TestFormatFromContentString_1(t *testing.T) {
	type testCase struct {
		name     string
		data     string
		expected Format
	}

	testCases := []testCase{
		{
			name:     "CSV",
			data:     "field1,field2,field3",
			expected: CSV,
		},
		{
			name:     "JSON",
			data:     "{\"field1\":\"value1\",\"field2\":\"value2\"}",
			expected: JSON,
		},
		{
			name:     "YAML",
			data:     "field1: value1\nfield2: value2",
			expected: YAML,
		},
		{
			name:     "XML",
			data:     "<root><field1>value1</field1><field2>value2</field2></root>",
			expected: XML,
		},
		{
			name:     "TOML",
			data:     "field1 = \"value1\"\nfield2 = \"value2\"",
			expected: TOML,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := Decoder{Delimiter: ','}
			result := d.FormatFromContentString(tc.data)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
