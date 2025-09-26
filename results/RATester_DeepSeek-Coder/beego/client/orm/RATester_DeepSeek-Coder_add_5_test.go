package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestAdd_5(t *testing.T) {
	type args struct {
		names []string
		mi    *models.ModelInfo
		fi    *models.FieldInfo
		inner bool
	}
	tests := []struct {
		name  string
		t     *dbTables
		args  args
		want  *dbTable
		want1 bool
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

			got, got1 := tt.t.add(tt.args.names, tt.args.mi, tt.args.fi, tt.args.inner)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbTables.add() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("dbTables.add() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
