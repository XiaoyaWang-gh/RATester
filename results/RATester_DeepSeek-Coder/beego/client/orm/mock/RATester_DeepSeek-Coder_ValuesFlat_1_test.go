package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValuesFlat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	result := &orm.ParamsList{}
	expr := "expr"

	count, err := d.ValuesFlat(result, expr)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if count != 0 {
		t.Errorf("Expected count to be 0, got %v", count)
	}

	if len(*result) != 0 {
		t.Errorf("Expected result to be empty, got %v", *result)
	}
}
