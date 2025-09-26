package redis

import (
	"fmt"
	"testing"
)

func TestStartAndGC_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rc := &Cache{
		conninfo: "localhost:6379",
		dbNum:    0,
		key:      "test",
		password: "",
		maxIdle:  10,
	}

	err := rc.StartAndGC("localhost:6379")
	if err != nil {
		t.Errorf("StartAndGC failed, err: %v", err)
	}
}
