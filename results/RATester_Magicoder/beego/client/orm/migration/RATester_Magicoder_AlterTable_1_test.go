package migration

import (
	"fmt"
	"testing"
)

func TestAlterTable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	tablename := "new_table"
	m.AlterTable(tablename)

	if m.TableName != tablename {
		t.Errorf("Expected TableName to be %s, but got %s", tablename, m.TableName)
	}

	if m.ModifyType != "alter" {
		t.Errorf("Expected ModifyType to be 'alter', but got %s", m.ModifyType)
	}
}
