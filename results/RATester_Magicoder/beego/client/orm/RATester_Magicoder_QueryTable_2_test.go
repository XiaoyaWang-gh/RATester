package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQueryTable_2(t *testing.T) {
	type args struct {
		ptrStructOrTableName interface{}
	}
	tests := []struct {
		name string
		f    *filterOrmDecorator
		args args
		want QuerySeter
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

			if got := tt.f.QueryTable(tt.args.ptrStructOrTableName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterOrmDecorator.QueryTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
