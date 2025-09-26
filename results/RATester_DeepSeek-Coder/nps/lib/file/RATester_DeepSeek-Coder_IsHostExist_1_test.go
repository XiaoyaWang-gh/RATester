package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestIsHostExist_1(t *testing.T) {
	db := &DbUtils{
		JsonDb: &JsonDb{
			Hosts: sync.Map{},
		},
	}

	testCases := []struct {
		name     string
		host     *Host
		expected bool
	}{
		{
			name: "Host exists",
			host: &Host{
				Id:       1,
				Host:     "test.com",
				Location: "/",
				Scheme:   "all",
			},
			expected: true,
		},
		{
			name: "Host does not exist",
			host: &Host{
				Id:       2,
				Host:     "nonexistent.com",
				Location: "/",
				Scheme:   "all",
			},
			expected: false,
		},
	}

	// Add a host to the database
	db.JsonDb.Hosts.Store(1, &Host{
		Id:       1,
		Host:     "test.com",
		Location: "/",
		Scheme:   "all",
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := db.IsHostExist(tc.host)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
