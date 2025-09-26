package context

import (
	"fmt"
	"testing"
)

func TestParam_2(t *testing.T) {
	type testStruct struct {
		name   string
		input  *BeegoInput
		key    string
		expect string
	}

	tests := []testStruct{
		{
			name: "TestParam_0",
			input: &BeegoInput{
				pnames:  []string{"key1", "key2", "key3"},
				pvalues: []string{"value1", "value2", "value3"},
			},
			key:    "key1",
			expect: "value1",
		},
		{
			name: "TestParam_1",
			input: &BeegoInput{
				pnames:  []string{"key1", "key2", "key3"},
				pvalues: []string{"value1", "value2", "value3"},
			},
			key:    "key2",
			expect: "value2",
		},
		{
			name: "TestParam_2",
			input: &BeegoInput{
				pnames:  []string{"key1", "key2", "key3"},
				pvalues: []string{"value1", "value2", "value3"},
			},
			key:    "key3",
			expect: "value3",
		},
		{
			name: "TestParam_3",
			input: &BeegoInput{
				pnames:  []string{"key1", "key2", "key3"},
				pvalues: []string{"value1", "value2", "value3"},
			},
			key:    "key4",
			expect: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.input.Param(test.key)
			if result != test.expect {
				t.Errorf("Expected %s, but got %s", test.expect, result)
			}
		})
	}
}
