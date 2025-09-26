package migration

import (
	"fmt"
	"testing"
)

func TestReset_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	m.sqls = append(m.sqls, "test")
	m.Reset()
	if len(m.sqls) != 0 {
		t.Errorf("sqls should be empty")
	}
}
