package order_clause

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestGetColumn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &Order{column: "id"}
	assert.Equal(t, "id", o.GetColumn())
}
