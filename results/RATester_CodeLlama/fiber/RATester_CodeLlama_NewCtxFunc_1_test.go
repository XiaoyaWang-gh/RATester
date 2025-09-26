package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestNewCtxFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	app := &App{}
	function := func(app *App) CustomCtx {
		return &DefaultCtx{}
	}
	// Act
	app.NewCtxFunc(function)
	// Assert
	assert.Equal(t, function, app.newCtxFunc)
}
