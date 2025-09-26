package mock

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestExclude_1(t *testing.T) {
	type args struct {
		s string
		i []interface{}
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

			if got := tt.d.Exclude(tt.args.s, tt.args.i...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exclude() = %v, want %v", got, tt.want)
			}
		})
	}
}
