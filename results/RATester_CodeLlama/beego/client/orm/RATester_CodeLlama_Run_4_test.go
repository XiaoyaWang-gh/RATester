package orm

import (
	"fmt"
	"testing"
)

func TestRun_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &commandSQLAll{}
	d.al = &alias{}
	err := d.Run()
	if err != nil {
		t.Errorf("Run() error = %v", err)
		return
	}
}
