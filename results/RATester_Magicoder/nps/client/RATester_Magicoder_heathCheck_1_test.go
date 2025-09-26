package client

import (
	"fmt"
	"testing"
	"time"

	"ehang.io/nps/lib/conn"
	"ehang.io/nps/lib/file"
)

func TestheathCheck_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	healths := []*file.Health{
		{
			HealthCheckTimeout:  10,
			HealthMaxFail:       5,
			HealthCheckInterval: 10,
			HealthNextTime:      time.Now(),
			HealthMap:           make(map[string]int),
			HttpHealthUrl:       "http://example.com",
			HealthRemoveArr:     []string{"remove1", "remove2"},
			HealthCheckType:     "type",
			HealthCheckTarget:   "target",
		},
	}

	c := &conn.Conn{
		Conn: nil,
		Rb:   []byte{},
	}

	result := heathCheck(healths, c)

	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}
