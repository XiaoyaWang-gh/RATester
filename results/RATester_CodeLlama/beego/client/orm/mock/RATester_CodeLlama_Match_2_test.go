package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestMatch_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	q := &QueryM2MerCondition{
		tableName: "tableName",
		name:      "name",
	}
	inv := &orm.Invocation{
		Args: []interface{}{"tableName", "name"},
	}
	if !q.Match(context.Background(), inv) {
		t.Errorf("q.Match(context.Background(), inv) = false, want true")
	}
}
