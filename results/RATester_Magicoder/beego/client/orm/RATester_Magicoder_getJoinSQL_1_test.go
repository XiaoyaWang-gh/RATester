package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestgetJoinSQL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dbTables := &dbTables{
		tablesM: map[string]*dbTable{},
		tables:  []*dbTable{},
		mi:      &models.ModelInfo{},
		base:    &dbBase{},
		skipEnd: false,
	}

	join := dbTables.getJoinSQL()

	if join == "" {
		t.Error("Expected a non-empty string, but got an empty string")
	}
}
