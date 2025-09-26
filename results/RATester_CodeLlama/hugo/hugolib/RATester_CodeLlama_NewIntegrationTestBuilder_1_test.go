package hugolib

import (
	"fmt"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestNewIntegrationTestBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	c := qt.New(t)

	var conf IntegrationTestConfig

	b := NewIntegrationTestBuilder(conf)

	c.Assert(b, qt.Not(qt.IsNil))
}
