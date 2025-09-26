package fiber

import (
	"fmt"
	"testing"
)

func TestBody_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		app: &App{
			config: Config{
				Immutable: true,
			},
		},
	}

	body := ctx.Body()

	if len(body) != 0 {
		t.Errorf("Expected empty body, got %s", body)
	}
}
