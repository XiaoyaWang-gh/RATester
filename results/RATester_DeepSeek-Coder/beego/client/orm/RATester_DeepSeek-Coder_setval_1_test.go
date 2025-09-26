package orm

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestSetval_1(t *testing.T) {
	type args struct {
		ctx        context.Context
		db         dbQuerier
		mi         *models.ModelInfo
		autoFields []string
	}
	tests := []struct {
		name    string
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

			d := &dbBasePostgres{}
			if err := d.setval(tt.args.ctx, tt.args.db, tt.args.mi, tt.args.autoFields); (err != nil) != tt.wantErr {
				t.Errorf("dbBasePostgres.setval() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
