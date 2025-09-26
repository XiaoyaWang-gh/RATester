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
	results := []orm.Params{}
	_, err := d.Values(&results)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
