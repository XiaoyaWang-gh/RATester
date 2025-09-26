package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestSet_33(t *testing.T) {
	type args struct {
		names []string
		mi    *models.ModelInfo
		fi    *models.FieldInfo
		inner bool
	}
	tests := []struct {
		name string
		t    *dbTables
		args args
		want *dbTable
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

			if got := tt.t.set(tt.args.names, tt.args.mi, tt.args.fi, tt.args.inner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbTables.set() = %v, want %v", got, tt.want)
			}
		})
	}
}
