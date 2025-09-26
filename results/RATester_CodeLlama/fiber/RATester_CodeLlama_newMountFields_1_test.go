package fiber

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewMountFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	app := &App{}

	// Act
	actual := newMountFields(app)

	// Assert
	assert.NotNil(t, actual)
	assert.Equal(t, app, actual.appList[""])
	assert.Equal(t, []string{""}, actual.appListKeys)
}
