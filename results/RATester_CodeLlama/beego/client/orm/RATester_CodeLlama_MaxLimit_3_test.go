package orm

import (
	"fmt"
	"testing"
)

func TestMaxLimit_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBasePostgres{}
	if d.MaxLimit() != 0 {
		t.Errorf("MaxLimit() = %v, want %v", d.MaxLimit(), 0)
	}
}
