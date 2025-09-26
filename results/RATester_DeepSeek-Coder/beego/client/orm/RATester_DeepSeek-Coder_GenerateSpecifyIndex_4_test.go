package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/hints"
)

func TestGenerateSpecifyIndex_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseSqlite{}
	tableName := "test_table"
	useIndex := hints.KeyUseIndex
	indexes := []string{"index1", "index2"}
	expected := fmt.Sprintf(` INDEXED BY %s,%s `, tableName+"`index1`", tableName+"`index2`")
	result := db.GenerateSpecifyIndex(tableName, useIndex, indexes)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
