package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValuesList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingRawSetter{}
	container := &[]orm.ParamsList{}
	cols := []string{"col1", "col2"}
	_, err := d.ValuesList(container, cols...)
	if err != nil {
		t.Errorf("ValuesList returned an error: %v", err)
	}
}
