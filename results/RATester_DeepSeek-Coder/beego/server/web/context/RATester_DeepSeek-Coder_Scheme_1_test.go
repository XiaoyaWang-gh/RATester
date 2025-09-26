package context

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestScheme_1(t *testing.T) {
	type testCase struct {
		name   string
		input  *BeegoInput
		expect string
	}

	testCases := []testCase{
		{
			name: "Test with X-Forwarded-Proto",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						Header: http.Header{
							"X-Forwarded-Proto": []string{"https"},
						},
					},
				},
			},
			expect: "https",
		},
		{
			name: "Test with Request.URL.Scheme",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						URL: &url.URL{
							Scheme: "http",
						},
					},
				},
			},
			expect: "http",
		},
		{
			name: "Test with Request.TLS",
			input: &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						TLS: &tls.ConnectionState{},
					},
				},
			},
			expect: "https",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.Scheme()
			if result != tc.expect {
				t.Errorf("Expected %s, got %s", tc.expect, result)
			}
		})
	}
}
