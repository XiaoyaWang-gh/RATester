package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
	"github.com/zeebo/assert"
)

func TestGetTableName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	inv := &Invocation{
		mi: &models.ModelInfo{
			Table: "table_name",
		},
	}
	assert.Equal(t, "table_name", inv.GetTableName())
}
