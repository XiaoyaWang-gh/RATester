package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestRowsToMap_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	d := &DoNothingRawSetter{}
	result := &orm.Params{}
	keyCol := ""
	valueCol := ""
	got, err := d.RowsToMap(result, keyCol, valueCol)
	if err != nil {
		t.Errorf("DoNothingRawSetter.RowsToMap() error = %v, wantErr %v", err, nil)
		return
	}
	if got != 0 {
		t.Errorf("DoNothingRawSetter.RowsToMap() = %v, want %v", got, 0)
	}
	// TODO: teardown test
}
