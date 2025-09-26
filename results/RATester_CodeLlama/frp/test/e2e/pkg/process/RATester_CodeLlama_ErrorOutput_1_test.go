package process

import (
	"bytes"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestErrorOutput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Process{}
	p.errorOutput = &bytes.Buffer{}
	p.errorOutput.WriteString("error output")
	assert.Equal(t, "error output", p.ErrorOutput())
}
