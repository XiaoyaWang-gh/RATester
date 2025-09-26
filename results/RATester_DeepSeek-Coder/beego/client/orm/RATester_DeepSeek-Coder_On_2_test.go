package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOn_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	cond := "a.b = c.d"
	expected := &PostgresQueryBuilder{
		tokens: []string{"ON", `"a"."b" = "c"."d"`},
	}
	qb.On(cond)
	if !reflect.DeepEqual(qb, expected) {
		t.Errorf("Expected %v, got %v", expected, qb)
	}
}
