package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestSetMaxOpenConns_1(t *testing.T) {
	type fields struct {
		Name            string
		Driver          DriverType
		DriverName      string
		DataSource      string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime time.Duration
		ConnMaxIdletime time.Duration
		StmtCacheSize   int
		DB              *DB
		DbBaser         dbBaser
		TZ              *time.Location
		Engine          string
	}
	type args struct {
		maxOpenConns int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			al := &alias{
				Name:            tt.fields.Name,
				Driver:          tt.fields.Driver,
				DriverName:      tt.fields.DriverName,
				DataSource:      tt.fields.DataSource,
				MaxIdleConns:    tt.fields.MaxIdleConns,
				MaxOpenConns:    tt.fields.MaxOpenConns,
				ConnMaxLifetime: tt.fields.ConnMaxLifetime,
				ConnMaxIdletime: tt.fields.ConnMaxIdletime,
				StmtCacheSize:   tt.fields.StmtCacheSize,
				DB:              tt.fields.DB,
				DbBaser:         tt.fields.DbBaser,
				TZ:              tt.fields.TZ,
				Engine:          tt.fields.Engine,
			}
			al.SetMaxOpenConns(tt.args.maxOpenConns)
		})
	}
}
