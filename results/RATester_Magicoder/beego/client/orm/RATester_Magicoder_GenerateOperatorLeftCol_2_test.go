package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestGenerateOperatorLeftCol_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseSqlite{}
	fi := &models.FieldInfo{
		FieldType: TypeDateField,
	}
	var leftCol string = "test_column"
	db.GenerateOperatorLeftCol(fi, "=", &leftCol)
	if fi.FieldType == TypeDateField {
		if leftCol != "DATE(test_column)" {
			t.Errorf("Expected DATE(test_column), got %s", leftCol)
		}
	}
}
