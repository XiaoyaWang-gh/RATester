package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestgetExistPk_1(t *testing.T) {
	type args struct {
		mi  *models.ModelInfo
		ind reflect.Value
	}
	tests := []struct {
		name       string
		args       args
		wantColumn string
		wantValue  interface{}
		wantExist  bool
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

			gotColumn, gotValue, gotExist := getExistPk(tt.args.mi, tt.args.ind)
			if gotColumn != tt.wantColumn {
				t.Errorf("getExistPk() gotColumn = %v, want %v", gotColumn, tt.wantColumn)
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("getExistPk() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotExist != tt.wantExist {
				t.Errorf("getExistPk() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}
