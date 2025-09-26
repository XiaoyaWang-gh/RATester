package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestRowsToMap_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{
		mi: &models.ModelInfo{},
	}

	result := &Params{}
	keyCol := "key"
	valueCol := "value"

	_, err := o.RowsToMap(result, keyCol, valueCol)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(*result) != 0 {
		t.Errorf("Expected empty result, got %v", *result)
	}
}
