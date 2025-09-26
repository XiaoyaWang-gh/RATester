package migration

import (
	"fmt"
	"testing"
)

func TestCreateTable_1(t *testing.T) {
	type args struct {
		tablename string
		engine    string
		charset   string
		p         []func()
	}
	tests := []struct {
		name string
		m    *Migration
		args args
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

			tt.m.CreateTable(tt.args.tablename, tt.args.engine, tt.args.charset, tt.args.p...)
		})
	}
}
