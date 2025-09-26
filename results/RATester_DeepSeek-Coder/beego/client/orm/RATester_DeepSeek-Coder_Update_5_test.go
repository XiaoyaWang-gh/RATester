package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestUpdate_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{
		mi: &models.ModelInfo{},
	}

	values := Params{
		"field1": "value1",
		"field2": "value2",
	}

	_, err := o.Update(values)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
