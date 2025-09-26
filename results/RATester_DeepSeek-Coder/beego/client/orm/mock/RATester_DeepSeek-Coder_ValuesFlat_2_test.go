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
	pl := &orm.ParamsList{}
	_, err := d.ValuesFlat(pl, "col1", "col2")
	if err != nil {
		t.Errorf("ValuesFlat() error = %v, wantErr %v", err, nil)
		return
	}
}
