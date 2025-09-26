package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestnewQuerySet_1(t *testing.T) {
	type args struct {
		orm *ormBase
		mi  *models.ModelInfo
	}
	tests := []struct {
		name string
		args args
		want QuerySeter
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

			if got := newQuerySet(tt.args.orm, tt.args.mi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQuerySet() = %v, want %v", got, tt.want)
			}
		})
	}
}
