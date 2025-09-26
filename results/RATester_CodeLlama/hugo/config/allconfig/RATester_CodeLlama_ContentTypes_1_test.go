package allconfig

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestContentTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	assert.Equal(t, c.ContentTypes(), c.config.C.ContentTypes)
}
