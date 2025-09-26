package lang

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/deps"
)

func TestTranslate_1(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ns := &Namespace{
		deps: &deps.Deps{},
	}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		id := "test_id"
		templateData := "test_data"
		expected := "translated_text"

		ns.deps.Translate = func(ctx context.Context, sid string, templateData any) string {
			assert.Equal(t, id, sid)
			assert.Equal(t, templateData, templateData)
			return expected
		}

		result, err := ns.Translate(ctx, id, templateData)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		id := "test_id"
		expectedErr := errors.New("test error")

		ns.deps.Translate = func(ctx context.Context, sid string, templateData any) string {
			assert.Equal(t, id, sid)
			return ""
		}

		_, err := ns.Translate(ctx, id, "test_data")
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("wrong number of arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		id := "test_id"
		_, err := ns.Translate(ctx, id, "test_data", "extra_arg")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "wrong number of arguments, expecting at most 2, got 3")
	})
}
