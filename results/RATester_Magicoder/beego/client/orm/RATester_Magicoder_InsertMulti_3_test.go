package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestInsertMulti_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	db := &dbBase{}
	mi := &models.ModelInfo{}
	sind := reflect.Value{}
	bulk := 10
	tz := time.Local

	_, err := db.InsertMulti(ctx, nil, mi, sind, bulk, tz)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
