package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestGenerateOperatorLeftCol_2(t *testing.T) {
	type args struct {
		fi       *models.FieldInfo
		operator string
		leftCol  *string
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

			d := &dbBaseSqlite{}
			d.GenerateOperatorLeftCol(tt.args.fi, tt.args.operator, tt.args.leftCol)
			if *tt.args.leftCol != tt.want {
				t.Errorf("GenerateOperatorLeftCol() = %v, want %v", *tt.args.leftCol, tt.want)
			}
		})
	}
}
