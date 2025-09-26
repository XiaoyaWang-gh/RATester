package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/hints"
)

func TestGenerateSpecifyIndex_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBase{}

	tableName := "test_table"
	useIndex := hints.KeyUseIndex
	indexes := []string{"index1", "index2"}

	expected := fmt.Sprintf(` USE INDEX(%s%s%s,%s%s%s) `, db.TableQuote(), indexes[0], db.TableQuote(), db.TableQuote(), indexes[1], db.TableQuote())
	result := db.GenerateSpecifyIndex(tableName, useIndex, indexes)

	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}
}
