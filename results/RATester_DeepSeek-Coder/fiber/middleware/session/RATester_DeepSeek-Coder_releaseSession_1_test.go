package session

import (
	"fmt"
	"testing"
)

func TestReleaseSession_1(t *testing.T) {
	t.Parallel()

	type args struct {
		s *Session
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "Test with a valid session",
			args: args{
				s: &Session{
					id: "test-id",
				},
			},
		},
		{
			name: "Test with a nil session",
			args: args{
				s: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			releaseSession(tc.args.s)
			if tc.args.s != nil {
				if tc.args.s.id != "" {
					t.Errorf("Expected session id to be empty, got %s", tc.args.s.id)
				}
				if tc.args.s.exp != 0 {
					t.Errorf("Expected session expiration to be 0, got %d", tc.args.s.exp)
				}
				if tc.args.s.config != nil {
					t.Errorf("Expected session config to be nil, got %v", tc.args.s.config)
				}
			}
		})
	}
}
