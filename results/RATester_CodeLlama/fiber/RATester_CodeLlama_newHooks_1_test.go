package fiber

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	app := &App{}
	// Act
	actual := newHooks(app)
	// Assert
	assert.NotNil(t, actual)
	assert.Equal(t, app, actual.app)
}
