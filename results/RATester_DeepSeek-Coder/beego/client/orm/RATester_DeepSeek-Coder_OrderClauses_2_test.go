package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/clauses/order_clause"
)

func TestOrderClauses_2(t *testing.T) {
	type args struct {
		orders []*order_clause.Order
	}
	tests := []struct {
		name string
		o    querySet
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

			if got := tt.o.OrderClauses(tt.args.orders...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("querySet.OrderClauses() = %v, want %v", got, tt.want)
			}
		})
	}
}
