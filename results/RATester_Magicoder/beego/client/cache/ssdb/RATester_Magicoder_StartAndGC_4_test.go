package ssdb

import (
	"fmt"
	"testing"
)

func TestStartAndGC_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{}
	config := `{"conn": "127.0.0.1:8888"}`
	err := cache.StartAndGC(config)
	if err != nil {
		t.Errorf("StartAndGC failed, err: %v", err)
	}
}
