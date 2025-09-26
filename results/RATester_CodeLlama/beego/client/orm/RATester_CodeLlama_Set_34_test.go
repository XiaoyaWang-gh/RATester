package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSet_34(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	kv := []string{"a", "b"}
	qb.Set(kv...)
	assert.Equal(t, "SET a=b", qb.String())
}
