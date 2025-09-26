package logger

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestBeforeHandlerFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	cfg := Config{}
	// WHEN
	beforeHandlerFunc(cfg)
	// THEN
	assert.NotNil(t, cfg.Output)
}
