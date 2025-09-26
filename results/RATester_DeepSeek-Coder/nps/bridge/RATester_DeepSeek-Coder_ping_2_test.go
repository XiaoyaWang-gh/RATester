package bridge

import (
	"fmt"
	"sync"
	"testing"
)

func TestPing_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bridge := &Bridge{
		Client: sync.Map{},
	}

	// Add some test data to the sync.Map
	bridge.Client.Store(1, &Client{
		retryTime: 0,
	})
	bridge.Client.Store(2, &Client{
		retryTime: 2,
	})
	bridge.Client.Store(3, &Client{
		retryTime: 3,
	})

	// Call the ping method
	bridge.ping()

	// Check the result
	bridge.Client.Range(func(key, value interface{}) bool {
		v := value.(*Client)
		if key.(int) == 1 || key.(int) == 2 || key.(int) == 3 {
			if v.retryTime != 1 {
				t.Errorf("Expected retryTime to be 1, got %d", v.retryTime)
			}
		} else {
			t.Errorf("Unexpected key %d", key)
		}
		return true
	})
}
