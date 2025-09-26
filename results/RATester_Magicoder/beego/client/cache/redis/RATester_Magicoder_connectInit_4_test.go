package redis

import (
	"fmt"
	"testing"
)

func TestconnectInit_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rc := &Cache{
		conninfo: "localhost:6379",
		dbNum:    0,
		maxIdle:  10,
	}
	rc.connectInit()

	if rc.p == nil {
		t.Error("Pool is nil")
	}

	if rc.p.Dial == nil {
		t.Error("Dial function is nil")
	}

	if rc.p.IdleTimeout != rc.timeout {
		t.Error("IdleTimeout is not set correctly")
	}

	if rc.p.MaxIdle != rc.maxIdle {
		t.Error("MaxIdle is not set correctly")
	}
}
