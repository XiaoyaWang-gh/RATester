package conn

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestGetLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	conn := &Conn{}
	// when
	_, err := conn.GetLen()
	// then
	assert.NoError(t, err)
}
