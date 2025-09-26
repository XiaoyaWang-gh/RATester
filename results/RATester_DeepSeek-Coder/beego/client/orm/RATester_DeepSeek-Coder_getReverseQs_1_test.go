package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestGetReverseQs_1(t *testing.T) {
	type args struct {
		md interface{}
		mi *models.ModelInfo
		fi *models.FieldInfo
	}
	tests := []struct {
		name string
		args args
		want *querySet
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

			o := &ormBase{}
			if got := o.getReverseQs(tt.args.md, tt.args.mi, tt.args.fi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReverseQs() = %v, want %v", got, tt.want)
			}
		})
	}
}
