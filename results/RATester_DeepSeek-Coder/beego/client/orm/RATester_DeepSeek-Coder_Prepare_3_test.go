package orm

import (
	"fmt"
	"testing"
)

func TestPrepare_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &rawSet{
		query: "SELECT * FROM test",
		args:  []interface{}{1, 2, 3},
	}

	rp, err := o.Prepare()
	if err != nil {
		t.Errorf("Prepare failed: %v", err)
		return
	}

	if rp == nil {
		t.Errorf("Prepare returned nil")
		return
	}

	_, err = rp.Exec(o.args...)
	if err != nil {
		t.Errorf("Exec failed: %v", err)
		return
	}

	err = rp.Close()
	if err != nil {
		t.Errorf("Close failed: %v", err)
		return
	}
}
