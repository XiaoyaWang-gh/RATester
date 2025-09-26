package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValuesList_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	results := &[]orm.ParamsList{}
	exprs := []string{"expr1", "expr2"}
	_, err := d.ValuesList(results, exprs...)
	if err != nil {
		t.Errorf("ValuesList() error = %v, wantErr %v", err, nil)
		return
	}
}
