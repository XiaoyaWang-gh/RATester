package fiber

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	assert.NotNil(t, app.Hooks())
}
