package mock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueryM2MerCondition_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tableName := "tableName"
	name := "name"
	q := NewQueryM2MerCondition(tableName, name)
	assert.NotNil(t, q)
	assert.Equal(t, tableName, q.tableName)
	assert.Equal(t, name, q.name)
}
