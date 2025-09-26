package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestSetMaxIdleConns_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	al := &alias{
		MaxIdleConns: 10,
		DB: &DB{
			DB: &sql.DB{},
		},
	}

	al.SetMaxIdleConns(20)

	if al.MaxIdleConns != 20 {
		t.Errorf("Expected MaxIdleConns to be 20, but got %d", al.MaxIdleConns)
	}

	if al.DB.DB.Stats().MaxIdleClosed != 20 {
		t.Errorf("Expected MaxIdleClosed to be 20, but got %d", al.DB.DB.Stats().MaxIdleClosed)
	}
}
