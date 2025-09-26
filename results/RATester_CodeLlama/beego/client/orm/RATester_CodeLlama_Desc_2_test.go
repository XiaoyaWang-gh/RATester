package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDesc_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Desc()
	assert.Equal(t, "DESC", qb.String())
}
