package dynacache

import (
	"fmt"
	"testing"
	"time"
)

func Teststart_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		opts: Options{
			CheckInterval: 1 * time.Second,
		},
	}

	stop := cache.start()
	defer stop()

	time.Sleep(2 * time.Second)

	// Add assertions here to check if the cache is working as expected
}
