package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestGetTableName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	inv := &Invocation{
		mi: &models.ModelInfo{
			Table: "test_table",
		},
	}

	tableName := inv.GetTableName()

	if tableName != "test_table" {
		t.Errorf("Expected table name to be 'test_table', but got '%s'", tableName)
	}
}
