package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestInsertOrUpdateSQL_1(t *testing.T) {
	type args struct {
		names  []string
		values *[]interface{}
		mi     *models.ModelInfo
		a      *alias
		args   []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
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

			d := &dbBase{}
			got, err := d.InsertOrUpdateSQL(tt.args.names, tt.args.values, tt.args.mi, tt.args.a, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertOrUpdateSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InsertOrUpdateSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
