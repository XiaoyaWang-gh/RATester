package mock

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestMockQueryM2MWithCtx_1(t *testing.T) {
	type args struct {
		tableName string
		name      string
		res       orm.QueryM2Mer
	}
	tests := []struct {
		name string
		args args
		want *Mock
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

			if got := MockQueryM2MWithCtx(tt.args.tableName, tt.args.name, tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockQueryM2MWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
