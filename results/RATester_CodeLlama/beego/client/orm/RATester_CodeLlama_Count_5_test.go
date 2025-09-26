package orm

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestCount_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d dbBase
	var q dbQuerier
	var qs querySet
	var mi *models.ModelInfo
	var cond *Condition
	var tz *time.Location
	var cnt int64
	var err error

	cnt, err = d.Count(context.Background(), q, qs, mi, cond, tz)
	if err != nil {
		t.Errorf("Count() error = %v", err)
		return
	}
	if cnt != 0 {
		t.Errorf("Count() = %v, want %v", cnt, 0)
	}
}
