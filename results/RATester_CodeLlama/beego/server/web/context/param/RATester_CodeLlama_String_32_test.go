package param

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestString_32(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mp := &MethodParam{
		name:         "name",
		in:           path,
		required:     true,
		defaultValue: "default",
	}
	assert.Equal(t, "param.New(\"name\", param.IsRequired, param.InPath, param.Default(\"default\"))", mp.String())
}
