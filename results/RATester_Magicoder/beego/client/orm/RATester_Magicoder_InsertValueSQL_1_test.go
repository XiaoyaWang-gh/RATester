package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestInsertValueSQL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBase{}

	names := []string{"name", "age"}
	values := []interface{}{"John", 30}
	isMulti := false
	mi := &models.ModelInfo{
		Table: "users",
	}

	query := db.InsertValueSQL(names, values, isMulti, mi)

	expected := "INSERT INTO `users` (`name`, `age`) VALUES (?, ?)"
	if query != expected {
		t.Errorf("Expected %s, got %s", expected, query)
	}
}
