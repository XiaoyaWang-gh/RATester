package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	d := &DoNothingQuerySetter{}
	results := &[]orm.Params{}
	exprs := []string{}
	got, err := d.Values(results, exprs...)
	if err != nil {
		t.Errorf("Values() error = %v", err)
		return
	}
	if got != 0 {
		t.Errorf("Values() got = %v, want %v", got, 0)
	}
}
