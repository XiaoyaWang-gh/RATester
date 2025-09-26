package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestInsertOrUpdate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	md := &models.ModelInfo{}
	colConflictAndArgs := []string{}
	_, err := o.InsertOrUpdate(md, colConflictAndArgs...)
	if err != nil {
		t.Errorf("TestInsertOrUpdate error, err:%v", err)
	}
}
