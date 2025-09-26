package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestRowsToStruct_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Key   string
		Value string
	}

	o := querySet{
		mi: &models.ModelInfo{},
	}

	_, err := o.RowsToStruct(&TestStruct{}, "key", "value")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
