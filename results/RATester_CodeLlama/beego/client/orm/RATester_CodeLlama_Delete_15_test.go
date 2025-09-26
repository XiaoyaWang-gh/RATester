package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Delete("table1")
	assert.Equal(t, "DELETE table1", qb.String())
}
