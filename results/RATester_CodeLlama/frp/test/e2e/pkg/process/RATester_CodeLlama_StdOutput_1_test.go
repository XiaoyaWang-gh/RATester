package process

import (
	"bytes"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestStdOutput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Process{}
	p.stdOutput = &bytes.Buffer{}
	p.stdOutput.WriteString("hello")
	assert.Equal(t, "hello", p.StdOutput())
}
