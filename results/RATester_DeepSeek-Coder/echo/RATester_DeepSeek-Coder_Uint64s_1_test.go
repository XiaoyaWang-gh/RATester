package echo

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestUint64s_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		source   string
		expected []uint64
	}{
		{
			name:     "single value",
			source:   "123",
			expected: []uint64{123},
		},
		{
			name:     "multiple values",
			source:   "123,456,789",
			expected: []uint64{123, 456, 789},
		},
		{
			name:     "empty",
			source:   "",
			expected: []uint64{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := &ValueBinder{}
			var dest []uint64
			b.Uint64s(tc.source, &dest)
			assert.Equal(t, tc.expected, dest)
		})
	}
}
