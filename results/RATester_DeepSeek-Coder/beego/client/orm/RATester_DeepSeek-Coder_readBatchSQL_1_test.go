package orm

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestReadBatchSQL_1(t *testing.T) {
	type args struct {
		tables *dbTables
		tCols  []string
		cond   *Condition
		qs     querySet
		mi     *models.ModelInfo
		tz     *time.Location
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []interface{}
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

			d := &dbBase{}
			got, got1 := d.readBatchSQL(tt.args.tables, tt.args.tCols, tt.args.cond, tt.args.qs, tt.args.mi, tt.args.tz)
			if got != tt.want {
				t.Errorf("readBatchSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("readBatchSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
