package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestNewQueryM2M_1(t *testing.T) {
	type args struct {
		md  interface{}
		o   *ormBase
		mi  *models.ModelInfo
		fi  *models.FieldInfo
		ind reflect.Value
	}
	tests := []struct {
		name string
		args args
		want QueryM2Mer
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

			got := newQueryM2M(tt.args.md, tt.args.o, tt.args.mi, tt.args.fi, tt.args.ind)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueryM2M() = %v, want %v", got, tt.want)
			}
		})
	}
}
