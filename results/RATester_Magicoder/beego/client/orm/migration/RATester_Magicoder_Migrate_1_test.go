package migration

import (
	"fmt"
	"testing"
)

func TestMigrate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}

	m.Migrate("test")

	if m.ModifyType != "test" {
		t.Errorf("Expected ModifyType to be 'test', but got '%s'", m.ModifyType)
	}

	if len(m.sqls) != 1 {
		t.Errorf("Expected sqls to have 1 item, but got %d", len(m.sqls))
	}
}
