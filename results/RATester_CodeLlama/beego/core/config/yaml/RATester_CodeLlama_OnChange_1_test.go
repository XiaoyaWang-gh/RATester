package yaml

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestOnChange_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given:
	var value string
	var configContainer ConfigContainer
	// when:
	configContainer.OnChange("", func(v string) {
		value = v
	})
	// then:
	assert.Equal(t, "", value)
}
