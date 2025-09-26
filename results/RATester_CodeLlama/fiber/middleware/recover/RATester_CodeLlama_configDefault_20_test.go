package recover //nolint:predeclared // TODO: Rename to some non-builtin
import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestConfigDefault_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := []Config{
		{
			EnableStackTrace: true,
		},
	}
	// Act
	actual := configDefault(config...)
	// Assert
	assert.Equal(t, actual.EnableStackTrace, true)
}
