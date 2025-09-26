package orm

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestsetColsValues_1(t *testing.T) {
	type args struct {
		mi     *models.ModelInfo
		ind    *reflect.Value
		cols   []string
		values []interface{}
		tz     *time.Location
	}
	tests := []struct {
		name string
		args args
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
			d.setColsValues(tt.args.mi, tt.args.ind, tt.args.cols, tt.args.values, tt.args.tz)
		})
	}
}
