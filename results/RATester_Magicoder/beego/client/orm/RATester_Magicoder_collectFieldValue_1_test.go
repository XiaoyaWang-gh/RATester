package orm

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestcollectFieldValue_1(t *testing.T) {
	type args struct {
		mi     *models.ModelInfo
		fi     *models.FieldInfo
		ind    reflect.Value
		insert bool
		tz     *time.Location
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
			got, err := d.collectFieldValue(tt.args.mi, tt.args.fi, tt.args.ind, tt.args.insert, tt.args.tz)
			if (err != nil) != tt.wantErr {
				t.Errorf("collectFieldValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collectFieldValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
