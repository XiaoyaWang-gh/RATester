package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOr_2(t *testing.T) {
	type args struct {
		expr string
		args []interface{}
	}
	tests := []struct {
		name string
		c    Condition
		args args
		want *Condition
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

			if got := tt.c.Or(tt.args.expr, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Condition.Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
