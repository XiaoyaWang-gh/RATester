package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestsetFieldValue_1(t *testing.T) {
	type args struct {
		fi    *models.FieldInfo
		value interface{}
		field reflect.Value
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
			got, err := d.setFieldValue(tt.args.fi, tt.args.value, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("setFieldValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setFieldValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
