package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeToDB_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	loc, _ := time.LoadLocation("UTC")
	now := time.Now()
	db := &dbBase{}
	db.TimeToDB(&now, loc)

	if now.Location() != loc {
		t.Errorf("Expected time to be in UTC, but it was in %s", now.Location())
	}
}
