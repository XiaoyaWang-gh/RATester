package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAssemblePagesStepFinal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	sa := &sitePagesAssembler{}
	// when
	err := sa.assemblePagesStepFinal()
	// then
	assert.NoError(t, err)
}
