package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFromDB_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	loc, _ := time.LoadLocation("UTC")
	db := &dbBase{}
	now := time.Now()
	db.TimeFromDB(&now, loc)

	if now.Location() != loc {
		t.Errorf("Expected time to be in UTC, but it was in %s", now.Location())
	}
}
