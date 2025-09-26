package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestSetConnMaxIdleTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	al := &alias{
		ConnMaxIdletime: 10 * time.Second,
	}

	al.SetConnMaxIdleTime(20 * time.Second)

	if al.ConnMaxIdletime != 20*time.Second {
		t.Errorf("Expected ConnMaxIdletime to be 20 seconds, but got %v", al.ConnMaxIdletime)
	}
}
