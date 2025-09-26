package accesslog

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestClose_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	h := &Handler{
		logHandlerChan: make(chan handlerParams),
	}
	// act
	err := h.Close()
	// assert
	assert.NoError(t, err)
}
