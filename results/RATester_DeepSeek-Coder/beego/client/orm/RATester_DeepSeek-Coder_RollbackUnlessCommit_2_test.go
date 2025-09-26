package orm

import (
	"fmt"
	"testing"
)

func TestRollbackUnlessCommit_2(t *testing.T) {
	type fields struct {
		alias *alias
		db    dbQuerier
		tx    txer
		txe   txEnder
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

			d := &dbQueryLog{
				alias: tt.fields.alias,
				db:    tt.fields.db,
				tx:    tt.fields.tx,
				txe:   tt.fields.txe,
			}
			if err := d.RollbackUnlessCommit(); (err != nil) != tt.wantErr {
				t.Errorf("dbQueryLog.RollbackUnlessCommit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
