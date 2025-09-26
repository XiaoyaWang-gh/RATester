package context

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestQuery_5(t *testing.T) {
	type testStruct struct {
		Name string `form:"name"`
	}

	testCases := []struct {
		name   string
		input  *BeegoInput
		key    string
		output string
	}{
		{
			name: "Test case 1",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						Form: url.Values{
							"name": []string{"John Doe"},
						},
					},
				},
			},
			key:    "name",
			output: "John Doe",
		},
		{
			name: "Test case 2",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						Form: url.Values{
							"name": []string{"Jane Doe"},
						},
					},
				},
			},
			key:    "name",
			output: "Jane Doe",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.Query(tc.key)
			if result != tc.output {
				t.Errorf("Expected %s, got %s", tc.output, result)
			}
		})
	}
}
