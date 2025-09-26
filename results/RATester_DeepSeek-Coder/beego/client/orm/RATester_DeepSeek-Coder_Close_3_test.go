package orm

import (
	"fmt"
	"testing"
)

func TestClose_3(t *testing.T) {
	type fields struct {
		alias *alias
		query string
		stmt  stmtQuerier
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &stmtQueryLog{
				alias: tt.fields.alias,
				query: tt.fields.query,
				stmt:  tt.fields.stmt,
			}
			if err := d.Close(); (err != nil) != tt.wantErr {
				t.Errorf("stmtQueryLog.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
