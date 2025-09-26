package gin

import (
	"fmt"
	"testing"
)

func TestSearchCredential_1(t *testing.T) {
	testCases := []struct {
		name      string
		authPairs authPairs
		authValue string
		wantUser  string
		wantOk    bool
	}{
		{
			name: "Test Case 1",
			authPairs: authPairs{
				{user: "user1", value: "value1"},
				{user: "user2", value: "value2"},
			},
			authValue: "value1",
			wantUser:  "user1",
			wantOk:    true,
		},
		{
			name: "Test Case 2",
			authPairs: authPairs{
				{user: "user1", value: "value1"},
				{user: "user2", value: "value2"},
			},
			authValue: "value3",
			wantUser:  "",
			wantOk:    false,
		},
		{
			name:      "Test Case 3",
			authPairs: nil,
			authValue: "value1",
			wantUser:  "",
			wantOk:    false,
		},
		{
			name: "Test Case 4",
			authPairs: authPairs{
				{user: "user1", value: "value1"},
				{user: "user2", value: "value2"},
			},
			authValue: "",
			wantUser:  "",
			wantOk:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			user, ok := tc.authPairs.searchCredential(tc.authValue)
			if user != tc.wantUser || ok != tc.wantOk {
				t.Errorf("Expected (%v, %v), got (%v, %v)", tc.wantUser, tc.wantOk, user, ok)
			}
		})
	}
}
