package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestUpdateBatchSQL_1(t *testing.T) {
	type args struct {
		mi             *models.ModelInfo
		cols           []string
		values         []interface{}
		specifyIndexes string
		join           string
		where          string
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
			if got := d.UpdateBatchSQL(tt.args.mi, tt.args.cols, tt.args.values, tt.args.specifyIndexes, tt.args.join, tt.args.where); got != tt.want {
				t.Errorf("UpdateBatchSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
