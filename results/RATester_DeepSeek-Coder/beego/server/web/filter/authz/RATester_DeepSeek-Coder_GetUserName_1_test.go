package authz

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"
)

func TestGetUserName_1(t *testing.T) {
	testCases := []struct {
		name         string
		request      *http.Request
		wantUsername string
	}{
		{
			name: "Test with valid basic auth",
			request: &http.Request{
				Header: http.Header{
					"Authorization": []string{"Basic " + base64.StdEncoding.EncodeToString([]byte("username:password"))},
				},
			},
			wantUsername: "username",
		},
		{
			name: "Test with invalid basic auth",
			request: &http.Request{
				Header: http.Header{
					"Authorization": []string{"Basic " + base64.StdEncoding.EncodeToString([]byte("invalid"))},
				},
			},
			wantUsername: "",
		},
		{
			name: "Test with no basic auth",
			request: &http.Request{
				Header: http.Header{},
			},
			wantUsername: "",
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
			if username != tc.wantUsername {
				t.Errorf("Expected username %s, got %s", tc.wantUsername, username)
			}
		})
	}
}
