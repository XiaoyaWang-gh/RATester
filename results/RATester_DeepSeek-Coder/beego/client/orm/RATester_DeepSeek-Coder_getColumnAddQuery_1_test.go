package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestGetColumnAddQuery_1(t *testing.T) {
	type args struct {
		al *alias
		fi *models.FieldInfo
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

			if got := getColumnAddQuery(tt.args.al, tt.args.fi); got != tt.want {
				t.Errorf("getColumnAddQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
