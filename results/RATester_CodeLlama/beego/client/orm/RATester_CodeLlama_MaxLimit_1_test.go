package orm

import (
	"fmt"
	"testing"
)

func TestMaxLimit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseSqlite{}
	if d.MaxLimit() != 9223372036854775807 {
		t.Errorf("MaxLimit() = %v, want %v", d.MaxLimit(), 9223372036854775807)
	}
}
