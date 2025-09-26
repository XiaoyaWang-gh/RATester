package orm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestDeleteBatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d dbBase
	var q dbQuerier
	var qs *querySet
	var mi *models.ModelInfo
	var cond *Condition
	var tz *time.Location

	d.DeleteBatch(context.Background(), q, qs, mi, cond, tz)
}
