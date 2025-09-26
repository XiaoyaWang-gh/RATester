package orm

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestGenerateOperatorSQL_1(t *testing.T) {
	type args struct {
		mi       *models.ModelInfo
		fi       *models.FieldInfo
		operator string
		args     []interface{}
		tz       *time.Location
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
			got, got1 := d.GenerateOperatorSQL(tt.args.mi, tt.args.fi, tt.args.operator, tt.args.args, tt.args.tz)
			if got != tt.want {
				t.Errorf("GenerateOperatorSQL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GenerateOperatorSQL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
