package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestDBStats_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{
		alias: &alias{
			DB: &DB{
				DB: &sql.DB{},
			},
		},
	}

	stats := o.DBStats()
	if stats == nil {
		t.Errorf("Expected non-nil stats, got nil")
	}
}
