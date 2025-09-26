package orm

import (
	"fmt"
	"testing"
)

func TestSet_34(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	kv := []string{"key1", "value1", "key2", "value2"}
	expected := "SET key1 = value1, key2 = value2"

	qb.Set(kv...)

	if qb.tokens[len(qb.tokens)-1] != expected {
		t.Errorf("Expected '%s', got '%s'", expected, qb.tokens[len(qb.tokens)-1])
	}
}
