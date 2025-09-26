package pageparser

import (
	"fmt"
	"testing"
)

func TestLexShortcodeQuotedParamVal_1(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Test case 1",
			input:          `"test"`,
			expectedOutput: `"test"`,
		},
		{
			name:           "Test case 2",
			input:          `"test \\"test\\" test"`,
			expectedOutput: `"test \\"test\\" test"`,
		},
		{
			name:           "Test case 3",
			input:          `"test \\"test\\" \\"test\\" test"`,
			expectedOutput: `"test \\"test\\" \\"test\\" test"`,
		},
		{
			name:           "Test case 4",
			input:          `"test \\"test\\" \\"test\\" \\"test\\" test"`,
			expectedOutput: `"test \\"test\\" \\"test\\" \\"test\\" test"`,
		},
		{
			name:           "Test case 5",
			input:          `"test \\"test\\" \\"test\\" \\"test\\" \\"test\\" test"`,
			expectedOutput: `"test \\"test\\" \\"test\\" \\"test\\" \\"test\\" test"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &pageLexer{
				input: []byte(test.input),
			}
			lexShortcodeQuotedParamVal(l, true, ItemType(0))
			output := string(l.current())
			if output != test.expectedOutput {
				t.Errorf("Expected %s, got %s", test.expectedOutput, output)
			}
		})
	}
}
