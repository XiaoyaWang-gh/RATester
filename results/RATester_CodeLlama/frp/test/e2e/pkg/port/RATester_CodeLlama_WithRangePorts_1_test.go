package port

import (
	"fmt"
	"testing"
)

func TestWithRangePorts_1(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		from int
		to   int
	}{
		{
			name: "from is 0",
			from: 0,
			to:   100,
		},
		{
			name: "to is 0",
			from: 100,
			to:   0,
		},
		{
			name: "from is greater than to",
			from: 100,
			to:   10,
		},
		{
			name: "from is less than to",
			from: 10,
			to:   100,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			builder := &nameBuilder{}
			WithRangePorts(tc.from, tc.to)(builder)
			if builder.rangePortFrom != tc.from {
				t.Errorf("WithRangePorts(%d, %d) builder.rangePortFrom = %d, want %d", tc.from, tc.to, builder.rangePortFrom, tc.from)
			}
			if builder.rangePortTo != tc.to {
				t.Errorf("WithRangePorts(%d, %d) builder.rangePortTo = %d, want %d", tc.from, tc.to, builder.rangePortTo, tc.to)
			}
		})
	}
}
