package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestNewAliasWithDb_1(t *testing.T) {
	type args struct {
		aliasName  string
		driverName string
		db         *sql.DB
		params     []DBOption
	}
	tests := []struct {
		name    string
		args    args
		want    *alias
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

			got, err := newAliasWithDb(tt.args.aliasName, tt.args.driverName, tt.args.db, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("newAliasWithDb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAliasWithDb() = %v, want %v", got, tt.want)
			}
		})
	}
}
