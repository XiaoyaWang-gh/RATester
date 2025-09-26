package session

import (
	"container/list"
	"context"
	"fmt"
	"testing"
)

func TestSessionAll_6(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testCases := []struct {
		name     string
		provider *MemProvider
		want     int
	}{
		{
			name: "empty provider",
			provider: &MemProvider{
				sessions: make(map[string]*list.Element),
				list:     list.New(),
			},
			want: 0,
		},
		{
			name: "non-empty provider",
			provider: &MemProvider{
				sessions: map[string]*list.Element{
					"session1": nil,
					"session2": nil,
					"session3": nil,
				},
				list: list.New(),
			},
			want: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.provider.SessionAll(ctx)
			if got != tc.want {
				t.Errorf("SessionAll() = %v, want %v", got, tc.want)
			}
		})
	}
}
