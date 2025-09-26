package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestAddAliasWthDB_1(t *testing.T) {
	type args struct {
		aliasName  string
		driverName string
		db         *sql.DB
		params     []DBOption
	}
	tests := []struct {
		name    string
		args    args
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

			if err := AddAliasWthDB(tt.args.aliasName, tt.args.driverName, tt.args.db, tt.args.params...); (err != nil) != tt.wantErr {
				t.Errorf("AddAliasWthDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
