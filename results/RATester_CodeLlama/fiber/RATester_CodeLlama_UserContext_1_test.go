package fiber

import (
	"context"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestUserContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{}
	ctx := c.UserContext()
	assert.Equal(t, context.Background(), ctx)
}
