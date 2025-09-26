package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestHasReturningID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d dbBasePostgres
	var mi *models.ModelInfo
	var query *string
	if got := d.HasReturningID(mi, query); got != false {
		t.Errorf("dbBasePostgres.HasReturningID() = %v, want %v", got, false)
	}
}
