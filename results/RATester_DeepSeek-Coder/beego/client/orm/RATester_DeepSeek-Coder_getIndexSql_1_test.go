package orm

import (
	"fmt"
	"testing"
)

func TestGetIndexSql_1(t *testing.T) {
	type args struct {
		tableName string
		useIndex  int
		indexes   []string
	}
	tests := []struct {
		name       string
		args       args
		wantClause string
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

			dt := &dbTables{}
			if gotClause := dt.getIndexSql(tt.args.tableName, tt.args.useIndex, tt.args.indexes); gotClause != tt.wantClause {
				t.Errorf("getIndexSql() = %v, want %v", gotClause, tt.wantClause)
			}
		})
	}
}
