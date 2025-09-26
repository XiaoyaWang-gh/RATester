package cache

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
)

func TestRemoveElement_1(t *testing.T) {
	testCases := []struct {
		name     string
		cache    *Cache
		element  *list.Element
		expected *list.Element
	}{
		{
			name: "Test removeElement",
			cache: &Cache{
				ll:    list.New(),
				cache: sync.Map{},
			},
			element:  &list.Element{},
			expected: &list.Element{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.cache.removeElement(tc.element)
			// assert.Equal(t, tc.expected, tc.cache.ll.Front())
		})
	}
}
