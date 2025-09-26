package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestHasReturningID_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBase{}
	mi := &models.ModelInfo{}
	id := ""
	result := db.HasReturningID(mi, &id)
	if result != false {
		t.Errorf("Expected false, but got %v", result)
	}
}
