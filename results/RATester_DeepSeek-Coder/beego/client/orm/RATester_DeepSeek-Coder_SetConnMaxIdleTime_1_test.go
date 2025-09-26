package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestSetConnMaxIdleTime_1(t *testing.T) {
	testCases := []struct {
		name     string
		idleTime time.Duration
	}{
		{
			name:     "Test case 1",
			idleTime: 1 * time.Minute,
		},
		{
			name:     "Test case 2",
			idleTime: 2 * time.Hour,
		},
		{
			name:     "Test case 3",
			idleTime: 3 * time.Second,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			al := &alias{
				ConnMaxIdletime: 0,
			}

			al.SetConnMaxIdleTime(tc.idleTime)

			if al.ConnMaxIdletime != tc.idleTime {
				t.Errorf("Expected ConnMaxIdletime to be %v, but got %v", tc.idleTime, al.ConnMaxIdletime)
			}
		})
	}
}
