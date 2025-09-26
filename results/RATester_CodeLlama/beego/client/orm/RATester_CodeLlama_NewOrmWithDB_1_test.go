package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestNewOrmWithDB_1(t *testing.T) {
	type args struct {
		driverName string
		aliasName  string
		db         *sql.DB
		params     []DBOption
	}
	tests := []struct {
		name    string
		args    args
		want    Ormer
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				driverName: "driverName",
				aliasName:  "aliasName",
				db:         &sql.DB{},
				params:     []DBOption{},
			},
			want:    &orm{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := NewOrmWithDB(tt.args.driverName, tt.args.aliasName, tt.args.db, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOrmWithDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrmWithDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}
