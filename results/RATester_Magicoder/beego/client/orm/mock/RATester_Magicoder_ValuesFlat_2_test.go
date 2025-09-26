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

	d := &DoNothingRawSetter{}
	container := &orm.ParamsList{}
	cols := []string{"col1", "col2"}
	_, err := d.ValuesFlat(container, cols...)
	if err != nil {
		t.Errorf("ValuesFlat returned an error: %v", err)
	}
}
