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
	result := &orm.Params{}
	keyCol := "key"
	valueCol := "value"
	_, err := d.RowsToMap(result, keyCol, valueCol)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
