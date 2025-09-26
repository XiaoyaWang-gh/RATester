package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/rs/zerolog"
)

func TestNewLogrusWrapper_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	logger := zerolog.New(os.Stdout)

	// act
	wrapper := NewLogrusWrapper(logger)

	// assert
	assert.NotNil(t, wrapper)
}
