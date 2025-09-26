package orm

import (
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/buffers"
	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestReadSQL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d dbBase
	var buf buffers.Buffer
	var tables *dbTables
	var tCols []string
	var cond *Condition
	var qs querySet
	var mi *models.ModelInfo
	var tz *time.Location
	d.readSQL(buf, tables, tCols, cond, qs, mi, tz)
}
