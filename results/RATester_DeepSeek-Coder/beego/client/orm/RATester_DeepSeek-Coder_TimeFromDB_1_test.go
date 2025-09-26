package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFromDB_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		time     time.Time
		location *time.Location
		expected time.Time
	}{
		{
			name:     "UTC",
			time:     time.Date(2022, 1, 1, 10, 0, 0, 0, time.UTC),
			location: time.UTC,
			expected: time.Date(2022, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "Local",
			time:     time.Date(2022, 1, 1, 10, 0, 0, 0, time.Local),
			location: time.Local,
			expected: time.Date(2022, 1, 1, 10, 0, 0, 0, time.Local),
		},
		{
			name:     "Empty Location",
			time:     time.Date(2022, 1, 1, 10, 0, 0, 0, time.UTC),
			location: nil,
			expected: time.Date(2022, 1, 1, 10, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			db := &dbBase{}
			db.TimeFromDB(&tc.time, tc.location)
			if !tc.time.Equal(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.time)
			}
		})
	}
}
