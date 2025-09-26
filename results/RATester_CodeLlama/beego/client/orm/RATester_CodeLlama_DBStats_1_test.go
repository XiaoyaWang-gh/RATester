package orm

import (
	"fmt"
	"testing"
)

func TestDBStats_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingOrm{}
	if got := d.DBStats(); got != nil {
		t.Errorf("DBStats() = %v, want nil", got)
	}
}
