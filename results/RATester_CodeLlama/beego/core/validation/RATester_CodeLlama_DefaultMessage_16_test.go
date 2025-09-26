package validation

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestDefaultMessage_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := Alpha{}
	assert.Equal(t, MessageTmpls["Alpha"], a.DefaultMessage())
}
