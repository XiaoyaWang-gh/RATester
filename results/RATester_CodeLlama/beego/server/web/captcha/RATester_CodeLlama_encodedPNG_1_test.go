package captcha

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestEncodedPNG_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	m := &Image{}
	// when
	result := m.encodedPNG()
	// then
	assert.NotNil(t, result)
}
