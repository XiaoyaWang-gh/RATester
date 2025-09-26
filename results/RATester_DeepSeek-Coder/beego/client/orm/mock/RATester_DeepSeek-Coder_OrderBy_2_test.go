package mock

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestOrderBy_2(t *testing.T) {
	t.Parallel()

	d := &DoNothingQuerySetter{}

	tests := []struct {
		name   string
		fields *DoNothingQuerySetter
		args   []string
		want   orm.QuerySeter
	}{
		{
			name:   "TestOrderBy",
			fields: d,
			args:   []string{"id"},
			want:   d,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.fields.OrderBy(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
