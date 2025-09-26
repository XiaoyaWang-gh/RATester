package orm

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"
)

func Testacquire_1(t *testing.T) {
	type fields struct {
		wg   sync.WaitGroup
		stmt *sql.Stmt
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test Case 1",
			fields: fields{
				wg:   sync.WaitGroup{},
				stmt: &sql.Stmt{},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &stmtDecorator{
				wg:   tt.fields.wg,
				stmt: tt.fields.stmt,
			}
			s.acquire()
		})
	}
}
