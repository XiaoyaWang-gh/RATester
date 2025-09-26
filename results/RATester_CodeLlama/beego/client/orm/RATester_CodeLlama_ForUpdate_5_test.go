package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForUpdate_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.ForUpdate()
	assert.Equal(t, "FOR UPDATE", qb.String())
}
