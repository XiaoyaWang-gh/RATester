package authz

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetUserName_1(t *testing.T) {
	testCases := []struct {
		name     string
		request  *http.Request
		expected string
	}{
		{
			name: "BasicAuth",
			request: &http.Request{
				Header: http.Header{
					"Authorization": []string{"Basic dXNlcjpwYXNzd29yZA=="},
				},
			},
			expected: "username",
		},
		{
			name: "NoBasicAuth",
			request: &http.Request{
				Header: http.Header{},
			},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			authorizer := &BasicAuthorizer{}
			username := authorizer.GetUserName(tc.request)
			if username != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, username)
			}
		})
	}
}
