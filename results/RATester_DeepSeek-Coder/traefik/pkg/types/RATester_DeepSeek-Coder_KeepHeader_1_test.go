package types

import (
	"fmt"
	"testing"
)

func TestKeepHeader_1(t *testing.T) {
	type testStruct struct {
		name     string
		fields   *AccessLogFields
		header   string
		expected string
	}

	tests := []testStruct{
		{
			name: "TestKeepHeader_HeaderExists",
			fields: &AccessLogFields{
				Headers: &FieldHeaders{
					DefaultMode: "drop",
					Names: map[string]string{
						"header1": "keep",
					},
				},
			},
			header:   "header1",
			expected: "keep",
		},
		{
			name: "TestKeepHeader_HeaderDoesNotExist",
			fields: &AccessLogFields{
				Headers: &FieldHeaders{
					DefaultMode: "drop",
					Names: map[string]string{
						"header2": "redact",
					},
				},
			},
			header:   "header1",
			expected: "drop",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.fields.KeepHeader(tt.header)
			if got != tt.expected {
				t.Errorf("KeepHeader() = %v, want %v", got, tt.expected)
			}
		})
	}
}
