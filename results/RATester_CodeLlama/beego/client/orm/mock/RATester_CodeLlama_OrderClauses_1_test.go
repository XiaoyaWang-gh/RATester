package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/clauses/order_clause"
)

func TestOrderClauses_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d *DoNothingQuerySetter
	var orders []*order_clause.Order
	d.OrderClauses(orders...)
}
