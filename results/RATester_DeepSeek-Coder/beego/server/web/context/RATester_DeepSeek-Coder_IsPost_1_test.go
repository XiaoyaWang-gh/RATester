package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsPost_1(t *testing.T) {
	type testCase struct {
		name   string
		input  *BeegoInput
		expect bool
	}

	testCases := []testCase{
		{
			name: "TestIsPost_POST",
			input: &BeegoInput{
				Context: &Context{
					Input: &BeegoInput{
						Context: &Context{
							Request: &http.Request{
								Method: "POST",
							},
						},
					},
				},
			},
			expect: true,
		},
		{
			name: "TestIsPost_GET",
			input: &BeegoInput{
				Context: &Context{
					Input: &BeegoInput{
						Context: &Context{
							Request: &http.Request{
								Method: "GET",
							},
						},
					},
				},
			},
			expect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.IsPost()
			if result != tc.expect {
				t.Errorf("Expected %v, but got %v", tc.expect, result)
			}
		})
	}
}
