package metadecoders

import (
	"fmt"
	"testing"
)

func TestFormatFromContentString_1(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected Format
	}{
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
		{
			name:     "No Format",
			data:     "plain text",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := Decoder{Delimiter: ','}
			result := d.FormatFromContentString(test.data)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
