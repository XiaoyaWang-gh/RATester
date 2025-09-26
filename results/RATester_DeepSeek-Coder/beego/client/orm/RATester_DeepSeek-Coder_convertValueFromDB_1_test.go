package orm

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestConvertValueFromDB_1(t *testing.T) {
	type args struct {
		fi  *models.FieldInfo
		val interface{}
		tz  *time.Location
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
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
			got, err := d.convertValueFromDB(tt.args.fi, tt.args.val, tt.args.tz)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertValueFromDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertValueFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
