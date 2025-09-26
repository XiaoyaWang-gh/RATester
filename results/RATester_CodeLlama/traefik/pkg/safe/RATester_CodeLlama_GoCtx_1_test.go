package safe

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGoCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Pool{}
	p.GoCtx(func(ctx context.Context) {
		assert.Equal(t, context.Background(), ctx)
	})
}
