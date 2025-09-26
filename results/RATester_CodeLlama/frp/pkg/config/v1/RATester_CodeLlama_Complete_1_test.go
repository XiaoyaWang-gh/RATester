package v1

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestComplete_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &ClientCommonConfig{
		User: "test",
	}
	c := &XTCPVisitorConfig{
		VisitorBaseConfig: VisitorBaseConfig{
			Name: "test",
		},
	}
	c.Complete(g)
	assert.Equal(t, "test", c.Name)
	assert.Equal(t, "test.test", c.FallbackTo)
}
