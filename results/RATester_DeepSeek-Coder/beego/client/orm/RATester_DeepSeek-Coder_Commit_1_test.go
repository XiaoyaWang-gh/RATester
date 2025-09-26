package orm

import (
	"fmt"
	"testing"
)

func TestCommit_1(t *testing.T) {
	type fields struct {
		ormBase ormBase
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

			tx := &txOrm{
				ormBase: tt.fields.ormBase,
			}
			if err := tx.Commit(); (err != nil) != tt.wantErr {
				t.Errorf("txOrm.Commit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
