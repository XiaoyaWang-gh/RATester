package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValuesFlat_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	d := &DoNothingRawSetter{}
	container := &orm.ParamsList{}
	cols := []string{}
	got, err := d.ValuesFlat(container, cols...)
	if err != nil {
		t.Errorf("ValuesFlat() error = %v", err)
		return
	}
	if got != 0 {
		t.Errorf("ValuesFlat() got = %v, want %v", got, 0)
	}
	// TODO: teardown test
}
