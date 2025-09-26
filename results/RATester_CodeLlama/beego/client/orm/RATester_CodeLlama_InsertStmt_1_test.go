package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestInsertStmt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d dbBase
	var stmt stmtQuerier
	var mi *models.ModelInfo
	var ind reflect.Value
	var tz *time.Location

	d.InsertStmt(context.Background(), stmt, mi, ind, tz)
}
