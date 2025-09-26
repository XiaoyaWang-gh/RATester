package postgres

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestconnectInit_2(t *testing.T) {
	type fields struct {
		maxlifetime int64
		savePath    string
	}
	tests := []struct {
		name   string
		fields fields
		want   *sql.DB
	}{
		{
			name: "Test Case 1",
			fields: fields{
				maxlifetime: 10,
				savePath:    "postgres://user:password@localhost/dbname?sslmode=disable",
			},
			want: func() *sql.DB {
				db, err := sql.Open("postgres", "postgres://user:password@localhost/dbname?sslmode=disable")
				if err != nil {
					t.Errorf("Error opening database: %v", err)
				}
				return db
			}(),
		},
		{
			name: "Test Case 2",
			fields: fields{
				maxlifetime: 20,
				savePath:    "postgres://user:password@localhost/dbname?sslmode=disable",
			},
			want: func() *sql.DB {
				db, err := sql.Open("postgres", "postgres://user:password@localhost/dbname?sslmode=disable")
				if err != nil {
					t.Errorf("Error opening database: %v", err)
				}
				return db
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			mp := &Provider{
				maxlifetime: tt.fields.maxlifetime,
				savePath:    tt.fields.savePath,
			}
			if got := mp.connectInit(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.connectInit() = %v, want %v", got, tt.want)
			}
		})
	}
}
