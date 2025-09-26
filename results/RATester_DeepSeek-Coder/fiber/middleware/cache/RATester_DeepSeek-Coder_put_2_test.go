package cache

import (
	"fmt"
	"testing"
)

func TestPut_2(t *testing.T) {
	h := &indexedHeap{
		entries: []heapEntry{},
		indices: []int{},
		maxidx:  0,
	}

	testCases := []struct {
		name   string
		key    string
		exp    uint64
		bytes  uint
		expect int
	}{
		{
			name:   "put first entry",
			key:    "key1",
			exp:    100,
			bytes:  10,
			expect: 0,
		},
		{
			name:   "put second entry",
			key:    "key2",
			exp:    200,
			bytes:  20,
			expect: 1,
		},
		{
			name:   "put third entry",
			key:    "key3",
			exp:    300,
			bytes:  30,
			expect: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := h.put(tc.key, tc.exp, tc.bytes)
			if result != tc.expect {
				t.Errorf("Expected %d, got %d", tc.expect, result)
			}
		})
	}
}
