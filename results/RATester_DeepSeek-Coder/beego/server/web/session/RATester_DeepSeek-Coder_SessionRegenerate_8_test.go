package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionRegenerate_8(t *testing.T) {
	ctx := context.Background()
	pder := &CookieProvider{}

	t.Run("Regenerate", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		store, err := pder.SessionRegenerate(ctx, "oldsid", "sid")
		if err != nil {
			t.Errorf("SessionRegenerate() error = %v", err)
			return
		}
		if store == nil {
			t.Errorf("SessionRegenerate() store = %v", store)
		}
	})
}
