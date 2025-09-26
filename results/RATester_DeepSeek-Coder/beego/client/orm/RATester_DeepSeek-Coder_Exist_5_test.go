package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/clauses/order_clause"
	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestExist_5(t *testing.T) {
	type fields struct {
		mi        *models.ModelInfo
		cond      *Condition
		related   []string
		relDepth  int
		limit     int64
		offset    int64
		groups    []string
		orders    []*order_clause.Order
		distinct  bool
		forUpdate bool
		useIndex  int
		indexes   []string
		orm       *ormBase
		aggregate string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
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

			o := querySet{
				mi:        tt.fields.mi,
				cond:      tt.fields.cond,
				related:   tt.fields.related,
				relDepth:  tt.fields.relDepth,
				limit:     tt.fields.limit,
				offset:    tt.fields.offset,
				groups:    tt.fields.groups,
				orders:    tt.fields.orders,
				distinct:  tt.fields.distinct,
				forUpdate: tt.fields.forUpdate,
				useIndex:  tt.fields.useIndex,
				indexes:   tt.fields.indexes,
				orm:       tt.fields.orm,
				aggregate: tt.fields.aggregate,
			}
			if got := o.Exist(); got != tt.want {
				t.Errorf("querySet.Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
