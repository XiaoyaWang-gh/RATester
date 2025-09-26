package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestOneWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given:
	d := &DoNothingQuerySetter{}
	ctx := context.Background()
	container := &struct{}{}
	cols := []string{"col1", "col2"}
	// when:
	err := d.OneWithCtx(ctx, container, cols...)
	// then:
	assert.NoError(t, err)
}
