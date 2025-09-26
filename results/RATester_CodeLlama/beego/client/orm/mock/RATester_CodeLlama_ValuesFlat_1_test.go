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

	// TODO: setup test
	d := &DoNothingQuerySetter{}
	result := &orm.ParamsList{}
	expr := ""
	// TODO: test and record result
	_, _ = d.ValuesFlat(result, expr)
	// TODO: verify results
}
