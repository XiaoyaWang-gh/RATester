package fiber

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestExecuteOnMountHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	app := &App{}
	h := &Hooks{}
	h.onMount = []OnMountHandler{
		func(app *App) error {
			return nil
		},
		func(app *App) error {
			return nil
		},
	}
	// Act
	err := h.executeOnMountHooks(app)
	// Assert
	assert.NoError(t, err)
}
