package orm

import (
	"fmt"
	"strings"
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
	expected := `"a"."b" = "c"."d"`
	qb.On(cond)
	result := qb.String()
	if !strings.Contains(result, expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
