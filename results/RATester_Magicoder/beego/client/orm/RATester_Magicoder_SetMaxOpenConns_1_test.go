package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestSetMaxOpenConns_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	al := &alias{
		MaxOpenConns: 10,
		DB: &DB{
			DB: &sql.DB{},
		},
	}

	al.SetMaxOpenConns(20)

	if al.MaxOpenConns != 20 {
		t.Errorf("Expected MaxOpenConns to be 20, but got %d", al.MaxOpenConns)
	}

	if al.DB.DB.Stats().MaxOpenConnections != 20 {
		t.Errorf("Expected MaxOpenConnections in DB to be 20, but got %d", al.DB.DB.Stats().MaxOpenConnections)
	}
}
