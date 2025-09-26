package migration

import (
	"fmt"
	"testing"
)

func TestExec_8(t *testing.T) {
	type args struct {
		name   string
		status string
	}
	tests := []struct {
		name    string
		m       *Migration
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

			m := &Migration{
				sqls:           tt.m.sqls,
				Created:        tt.m.Created,
				TableName:      tt.m.TableName,
				Engine:         tt.m.Engine,
				Charset:        tt.m.Charset,
				ModifyType:     tt.m.ModifyType,
				Columns:        tt.m.Columns,
				Indexes:        tt.m.Indexes,
				Primary:        tt.m.Primary,
				Uniques:        tt.m.Uniques,
				Foreigns:       tt.m.Foreigns,
				Renames:        tt.m.Renames,
				RemoveColumns:  tt.m.RemoveColumns,
				RemoveIndexes:  tt.m.RemoveIndexes,
				RemoveUniques:  tt.m.RemoveUniques,
				RemoveForeigns: tt.m.RemoveForeigns,
			}
			if err := m.Exec(tt.args.name, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Migration.Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
