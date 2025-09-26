package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestGetExistPk_1(t *testing.T) {
	type args struct {
		mi  *models.ModelInfo
		ind reflect.Value
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 interface{}
		want2 bool
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

			column, value, exist := getExistPk(tt.args.mi, tt.args.ind)
			if column != tt.want {
				t.Errorf("getExistPk() got = %v, want %v", column, tt.want)
			}
			if !reflect.DeepEqual(value, tt.want1) {
				t.Errorf("getExistPk() got1 = %v, want %v", value, tt.want1)
			}
			if exist != tt.want2 {
				t.Errorf("getExistPk() got2 = %v, want %v", exist, tt.want2)
			}
		})
	}
}
