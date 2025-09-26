package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func Testget_24(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testTables := &dbTables{
		tablesM: map[string]*dbTable{
			"testTable": {
				id:    1,
				index: "testIndex",
				name:  "testTable",
				names: []string{"testTable"},
				sel:   true,
				inner: false,
				mi:    &models.ModelInfo{},
				fi:    &models.FieldInfo{},
				jtl:   &dbTable{},
			},
		},
	}

	testName := "testTable"
	expectedTable := &dbTable{
		id:    1,
		index: "testIndex",
		name:  "testTable",
		names: []string{"testTable"},
		sel:   true,
		inner: false,
		mi:    &models.ModelInfo{},
		fi:    &models.FieldInfo{},
		jtl:   &dbTable{},
	}

	table, ok := testTables.get(testName)

	if !ok {
		t.Errorf("Expected table %s not found", testName)
	}

	if table.id != expectedTable.id ||
		table.index != expectedTable.index ||
		table.name != expectedTable.name ||
		table.sel != expectedTable.sel ||
		table.inner != expectedTable.inner ||
		table.mi != expectedTable.mi ||
		table.fi != expectedTable.fi ||
		table.jtl != expectedTable.jtl {
		t.Errorf("Expected table %v, got %v", expectedTable, table)
	}
}
