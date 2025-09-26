package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestInsertValueSQL_1(t *testing.T) {
	type args struct {
		names   []string
		values  []interface{}
		isMulti bool
		mi      *models.ModelInfo
	}
	tests := []struct {
		name string
		args args
		want string
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
			if got := d.InsertValueSQL(tt.args.names, tt.args.values, tt.args.isMulti, tt.args.mi); got != tt.want {
				t.Errorf("InsertValueSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
