package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestValuesFlat_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{
		mi: &models.ModelInfo{},
		cond: &Condition{
			params: []condValue{
				{
					exprs: []string{"id", "=", "1"},
				},
			},
		},
	}

	result := &ParamsList{}
	_, err := o.ValuesFlat(result, "id")
	if err != nil {
		t.Errorf("ValuesFlat() error = %v", err)
		return
	}

	if len(*result) != 1 {
		t.Errorf("ValuesFlat() result length = %v, want %v", len(*result), 1)
	}
}
