package orm

import (
	"fmt"
	"reflect"
	"testing"

	imodels "github.com/beego/beego/v2/client/orm/internal/models"
)

func TestGetDbDropSQL_1(t *testing.T) {
	type args struct {
		mc *imodels.ModelCache
		al *alias
	}
	tests := []struct {
		name    string
		args    args
		want    []string
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

			got, err := getDbDropSQL(tt.args.mc, tt.args.al)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDbDropSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDbDropSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
