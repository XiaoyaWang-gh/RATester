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
	expectedCount := int64(0)
	count, err := d.ValuesFlat(result, expr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if count != expectedCount {
		t.Errorf("Expected count %d, but got %d", expectedCount, count)
	}
}
