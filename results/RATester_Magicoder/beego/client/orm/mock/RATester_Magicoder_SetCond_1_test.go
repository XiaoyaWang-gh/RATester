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

	d := &DoNothingQuerySetter{}
	condition := orm.NewCondition()
	condition = condition.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)
	d.SetCond(condition)

	// Add your assertions here
}
