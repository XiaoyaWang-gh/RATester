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

	m := &Migration{
		sqls: []string{"sql1", "sql2"},
	}

	m.Reset()

	if len(m.sqls) != 0 {
		t.Errorf("Expected sqls to be empty after Reset, but got %v", m.sqls)
	}
}
