package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestSetCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d *DoNothingQuerySetter
	var condition *orm.Condition
	d.SetCond(condition)
}
