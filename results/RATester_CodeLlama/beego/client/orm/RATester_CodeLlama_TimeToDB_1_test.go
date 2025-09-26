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

	d := &dbBase{}
	t0 := time.Now()
	tz := time.FixedZone("test", 8*60*60)
	d.TimeToDB(&t0, tz)
	if t0.Location() != tz {
		t.Errorf("TimeToDB failed")
	}
}
