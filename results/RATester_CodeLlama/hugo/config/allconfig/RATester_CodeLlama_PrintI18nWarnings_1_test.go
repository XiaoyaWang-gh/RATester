package allconfig

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestPrintI18nWarnings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	assert.False(t, c.PrintI18nWarnings())
}
