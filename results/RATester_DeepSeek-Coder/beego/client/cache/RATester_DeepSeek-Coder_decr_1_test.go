package cache

import (
	"fmt"
	"math"
	"testing"
)

func TestDecr_1(t *testing.T) {
	type testCase struct {
		name     string
		input    interface{}
		expected interface{}
		err      error
	}

	testCases := []testCase{
		{"Test int overflow", math.MaxInt, nil, ErrDecrementOverflow},
		{"Test int32 overflow", math.MinInt32, nil, ErrDecrementOverflow},
		{"Test int64 overflow", math.MinInt64, nil, ErrDecrementOverflow},
		{"Test uint overflow", uint(0), nil, ErrDecrementOverflow},
		{"Test uint32 overflow", uint32(0), nil, ErrDecrementOverflow},
		{"Test uint64 overflow", uint64(0), nil, ErrDecrementOverflow},
		{"Test int decrement", 1, 0, nil},
		{"Test int32 decrement", int32(1), int32(0), nil},
		{"Test int64 decrement", int64(1), int64(0), nil},
		{"Test uint decrement", uint(1), uint(0), nil},
		{"Test uint32 decrement", uint32(1), uint32(0), nil},
		{"Test uint64 decrement", uint64(1), uint64(0), nil},
		{"Test non-integer type", "1", nil, ErrNotIntegerType},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := decr(tc.input)
			if err != tc.err {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
			if result != tc.expected {
				t.Errorf("Expected result %v, got %v", tc.expected, result)
			}
		})
	}
}
