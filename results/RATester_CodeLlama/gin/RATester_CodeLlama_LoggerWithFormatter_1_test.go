package gin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerWithFormatter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	f := func(params LogFormatterParams) string {
		return "test"
	}
	// when
	h := LoggerWithFormatter(f)
	// then
	assert.NotNil(t, h)
}
