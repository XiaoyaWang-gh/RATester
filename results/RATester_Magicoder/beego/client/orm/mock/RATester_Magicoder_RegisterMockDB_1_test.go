package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/orm"
)

func TestRegisterMockDB_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	RegisterMockDB(name)

	db, err := orm.GetDB(name)
	if err != nil {
		t.Errorf("Failed to get database: %v", err)
	}

	if db == nil {
		t.Error("Database is nil")
	}
}
