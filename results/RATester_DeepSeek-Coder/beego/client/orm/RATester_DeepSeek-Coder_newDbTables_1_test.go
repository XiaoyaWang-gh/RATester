package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestNewDbTables_1(t *testing.T) {
	type args struct {
		mi   *models.ModelInfo
		base dbBaser
	}
	tests := []struct {
		name string
		args args
		want *dbTables
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

			if got := newDbTables(tt.args.mi, tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDbTables() = %v, want %v", got, tt.want)
			}
		})
	}
}
