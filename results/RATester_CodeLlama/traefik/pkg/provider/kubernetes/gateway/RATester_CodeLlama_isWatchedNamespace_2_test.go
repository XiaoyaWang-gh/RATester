package gateway

import (
	"fmt"
	"testing"
)

func TestIsWatchedNamespace_2(t *testing.T) {
	tests := []struct {
		name string
		c    *clientWrapper
		want bool
	}{
		{
			name: "isWatchedNamespace",
			c: &clientWrapper{
				isNamespaceAll: true,
			},
			want: true,
		},
		{
			name: "isWatchedNamespace",
			c: &clientWrapper{
				isNamespaceAll:    false,
				watchedNamespaces: []string{"test"},
			},
			want: true,
		},
		{
			name: "isWatchedNamespace",
			c: &clientWrapper{
				isNamespaceAll:    false,
				watchedNamespaces: []string{"test1", "test2"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.isWatchedNamespace("test"); got != tt.want {
				t.Errorf("clientWrapper.isWatchedNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
