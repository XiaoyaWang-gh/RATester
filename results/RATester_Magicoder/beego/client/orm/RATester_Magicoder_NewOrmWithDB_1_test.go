package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestNewOrmWithDB_1(t *testing.T) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

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
			name: "Test case 1",
			args: args{
				driverName: "mysql",
				aliasName:  "test",
				db:         db,
				params:     []DBOption{},
			},
			want:    nil,
			wantErr: false,
		},
		// Add more test cases here
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
				t.Errorf("NewOrmWithDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
