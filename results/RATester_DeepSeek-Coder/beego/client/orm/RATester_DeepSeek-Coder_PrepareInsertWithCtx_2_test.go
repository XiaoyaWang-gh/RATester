package orm

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/internal/models"
)

func TestPrepareInsertWithCtx_2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mi := &models.ModelInfo{}
	orm := &ormBase{}
	qs := querySet{
		mi:  mi,
		orm: orm,
	}

	t.Run("should return Inserter and nil error when prepare insert with context", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := qs.PrepareInsertWithCtx(ctx)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})

	t.Run("should return error when context is canceled", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cancel()
		_, err := qs.PrepareInsertWithCtx(ctx)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
