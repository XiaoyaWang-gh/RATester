package orm

import (
	"fmt"
	"testing"
)

func TestDesc_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Desc()
	expected := "DESC"
	actual := qb.tokens[len(qb.tokens)-1]
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
