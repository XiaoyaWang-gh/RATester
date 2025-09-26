package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValues_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingRawSetter{}
	container := &[]orm.Params{}
	cols := []string{"col1", "col2"}
	got, err := d.Values(container, cols...)
	if err != nil {
		t.Fatalf("Values() error = %v", err)
	}
	if got != 0 {
		t.Errorf("Values() got = %v, want %v", got, 0)
	}
}
