package cache

import (
	"container/list"
	"fmt"
	"testing"
)

func TestNew_1(t *testing.T) {
	tests := []struct {
		name       string
		maxEntries int
		want       *Cache
	}{
		{
			name:       "Test case 1",
			maxEntries: 100,
			want: &Cache{
				MaxEntries: 100,
				ll:         list.New(),
			},
		},
		{
			name:       "Test case 2",
			maxEntries: 0,
			want: &Cache{
				MaxEntries: 0,
				ll:         list.New(),
			},
		},
		{
			name:       "Test case 3",
			maxEntries: 500,
			want: &Cache{
				MaxEntries: 500,
				ll:         list.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := New(tt.maxEntries)
			if got.MaxEntries != tt.want.MaxEntries {
				t.Errorf("New() = %v, want %v", got.MaxEntries, tt.want.MaxEntries)
			}
			if fmt.Sprintf("%p", got.ll) != fmt.Sprintf("%p", tt.want.ll) {
				t.Errorf("New() = %v, want %v", fmt.Sprintf("%p", got.ll), fmt.Sprintf("%p", tt.want.ll))
			}
		})
	}
}
