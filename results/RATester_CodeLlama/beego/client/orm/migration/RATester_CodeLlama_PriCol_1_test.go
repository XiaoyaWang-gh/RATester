package migration

import (
	"fmt"
	"testing"
)

func TestPriCol_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	m.PriCol("name")
	if len(m.Primary) != 1 {
		t.Errorf("PriCol failed")
	}
}
