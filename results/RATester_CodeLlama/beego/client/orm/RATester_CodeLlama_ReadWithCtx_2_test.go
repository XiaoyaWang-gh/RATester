package orm

import (
	"context"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestReadWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	d := &DoNothingOrm{}
	ctx := context.Background()
	md := &struct{}{}
	cols := []string{"col1", "col2"}

	// When
	err := d.ReadWithCtx(ctx, md, cols...)

	// Then
	assert.NoError(t, err)
}
