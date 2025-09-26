package orm

import (
	"fmt"
	"testing"
)

func TestName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := driver("mysql")
	if d.Name() != "mysql" {
		t.Errorf("driver.Name() = %v, want %v", d.Name(), "mysql")
	}
}
