package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrepareInsertWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	d := &DoNothingQuerySetter{}
	ctx := context.TODO()
	got, err := d.PrepareInsertWithCtx(ctx)
	require.NoError(t, err)
	require.NotNil(t, got)
	// TODO: teardown test
}
