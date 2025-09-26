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

	var d dbBaseSqlite
	var fi models.FieldInfo
	var operator string
	var leftCol string
	d.GenerateOperatorLeftCol(&fi, operator, &leftCol)
	if leftCol != "" {
		t.Errorf("TestGenerateOperatorLeftCol failed")
	}
}
