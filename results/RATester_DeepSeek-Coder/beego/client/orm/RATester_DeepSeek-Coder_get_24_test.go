package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestGet_24(t *testing.T) {
	tables := &dbTables{
		tablesM: map[string]*dbTable{
			"test": {
				id:    1,
				index: "test",
				name:  "test",
				names: []string{"test"},
				sel:   true,
				inner: true,
				mi:    &models.ModelInfo{},
				fi:    &models.FieldInfo{},
				jtl:   &dbTable{},
			},
		},
	}

	tests := []struct {
		name  string
		t     *dbTables
		arg   string
		want  *dbTable
		want1 bool
	}{
		{
			name:  "test",
			t:     tables,
			arg:   "test",
			want:  tables.tablesM["test"],
			want1: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := tt.t.get(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
