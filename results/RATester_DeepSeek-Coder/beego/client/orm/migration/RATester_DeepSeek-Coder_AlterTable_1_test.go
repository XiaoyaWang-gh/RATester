package migration

import (
	"fmt"
	"testing"
)

func TestAlterTable_1(t *testing.T) {
	type fields struct {
		TableName  string
		ModifyType string
	}
	type args struct {
		tablename string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "TestAlterTable",
			fields: fields{
				TableName:  "old_table",
				ModifyType: "create",
			},
			args: args{
				tablename: "new_table",
			},
			want: "new_table",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &Migration{
				TableName:  tt.fields.TableName,
				ModifyType: tt.fields.ModifyType,
			}
			m.AlterTable(tt.args.tablename)
			if m.TableName != tt.want {
				t.Errorf("AlterTable() = %v, want %v", m.TableName, tt.want)
			}
			if m.ModifyType != "alter" {
				t.Errorf("AlterTable() = %v, want %v", m.ModifyType, "alter")
			}
		})
	}
}
