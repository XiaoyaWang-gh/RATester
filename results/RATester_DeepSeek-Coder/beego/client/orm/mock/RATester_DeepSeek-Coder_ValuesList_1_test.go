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
	expectedCount := int64(0)
	expectedError := error(nil)

	count, err := d.ValuesList(container, cols...)

	if count != expectedCount {
		t.Errorf("Expected count to be %v, but got %v", expectedCount, count)
	}

	if err != expectedError {
		t.Errorf("Expected error to be %v, but got %v", expectedError, err)
	}
}
