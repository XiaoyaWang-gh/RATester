package context

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestIsSecure_1(t *testing.T) {
	type testCase struct {
		name   string
		input  *BeegoInput
		expect bool
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						URL: &url.URL{
							Scheme: "https",
						},
					},
				},
			},
			expect: true,
		},
		{
			name: "Test case 2",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						URL: &url.URL{
							Scheme: "http",
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

			result := tc.input.IsSecure()
			if result != tc.expect {
				t.Errorf("Expected %v, but got %v", tc.expect, result)
			}
		})
	}
}
