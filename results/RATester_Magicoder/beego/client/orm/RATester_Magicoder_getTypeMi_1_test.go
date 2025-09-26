package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestgetTypeMi_1(t *testing.T) {
	type args struct {
		mdTyp reflect.Type
	}
	tests := []struct {
		name string
		args args
		want *models.ModelInfo
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

			if got := getTypeMi(tt.args.mdTyp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTypeMi() = %v, want %v", got, tt.want)
			}
		})
	}
}
