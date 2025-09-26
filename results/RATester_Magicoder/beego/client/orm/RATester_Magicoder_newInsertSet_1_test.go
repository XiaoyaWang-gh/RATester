package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestnewInsertSet_1(t *testing.T) {
	type args struct {
		ctx context.Context
		orm *ormBase
		mi  *models.ModelInfo
	}
	tests := []struct {
		name    string
		args    args
		want    Inserter
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

			got, err := newInsertSet(tt.args.ctx, tt.args.orm, tt.args.mi)
			if (err != nil) != tt.wantErr {
				t.Errorf("newInsertSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInsertSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
