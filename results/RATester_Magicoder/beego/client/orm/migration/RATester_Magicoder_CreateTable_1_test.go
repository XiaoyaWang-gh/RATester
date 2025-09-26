package migration

import (
	"fmt"
	"testing"
)

func TestCreateTable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	tablename := "test_table"
	engine := "InnoDB"
	charset := "utf8"
	m.CreateTable(tablename, engine, charset)

	if m.TableName != tablename {
		t.Errorf("Expected TableName to be %s, but got %s", tablename, m.TableName)
	}

	if m.Engine != engine {
		t.Errorf("Expected Engine to be %s, but got %s", engine, m.Engine)
	}

	if m.Charset != charset {
		t.Errorf("Expected Charset to be %s, but got %s", charset, m.Charset)
	}

	if m.ModifyType != "create" {
		t.Errorf("Expected ModifyType to be 'create', but got %s", m.ModifyType)
	}
}
