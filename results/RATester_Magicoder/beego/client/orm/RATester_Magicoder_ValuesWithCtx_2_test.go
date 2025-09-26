package orm

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestValuesWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	o := &querySet{
		mi: &models.ModelInfo{},
	}
	results := make([]Params, 0)
	_, err := o.ValuesWithCtx(ctx, &results, "id", "name")
	if err != nil {
		t.Errorf("ValuesWithCtx failed, err: %v", err)
	}
}
