package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	var results []orm.Params
	_, err := d.Values(&results, "col1", "col2")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(results) != 0 {
		t.Errorf("Expected empty results, got %v", results)
	}
}
