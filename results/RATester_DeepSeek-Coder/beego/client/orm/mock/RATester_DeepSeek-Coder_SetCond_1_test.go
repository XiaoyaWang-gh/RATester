package mock

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestSetCond_1(t *testing.T) {
	type args struct {
		condition *orm.Condition
	}
	tests := []struct {
		name string
		d    *DoNothingQuerySetter
		args args
		want orm.QuerySeter
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

			if got := tt.d.SetCond(tt.args.condition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCond() = %v, want %v", got, tt.want)
			}
		})
	}
}
