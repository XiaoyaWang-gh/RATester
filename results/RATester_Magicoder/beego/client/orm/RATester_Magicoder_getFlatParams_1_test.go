package orm

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestgetFlatParams_1(t *testing.T) {
	type args struct {
		fi   *models.FieldInfo
		args []interface{}
		tz   *time.Location
	}
	tests := []struct {
		name       string
		args       args
		wantParams []interface{}
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

			if gotParams := getFlatParams(tt.args.fi, tt.args.args, tt.args.tz); !reflect.DeepEqual(gotParams, tt.wantParams) {
				t.Errorf("getFlatParams() = %v, want %v", gotParams, tt.wantParams)
			}
		})
	}
}
