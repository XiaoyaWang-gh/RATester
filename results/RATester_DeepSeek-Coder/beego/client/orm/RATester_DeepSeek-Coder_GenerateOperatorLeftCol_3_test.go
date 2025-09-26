package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestGenerateOperatorLeftCol_3(t *testing.T) {
	type args struct {
		fi  *models.FieldInfo
		op  string
		res *string
	}
	tests := []struct {
		name string
		args args
		want *string
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
			d.GenerateOperatorLeftCol(tt.args.fi, tt.args.op, tt.args.res)
			if !reflect.DeepEqual(tt.args.res, tt.want) {
				t.Errorf("GenerateOperatorLeftCol() = %v, want %v", tt.args.res, tt.want)
			}
		})
	}
}
