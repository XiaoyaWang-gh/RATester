package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestRowsToMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	result := make(orm.Params)
	keyCol := "key"
	valueCol := "value"
	count, err := d.RowsToMap(&result, keyCol, valueCol)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if count != 0 {
		t.Errorf("Expected count to be 0, but got %v", count)
	}
}
