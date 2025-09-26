package fiber

import (
	"fmt"
	"os"
	"testing"

	"github.com/zeebo/assert"
)

func TestIsChild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	os.Setenv(envPreforkChildKey, envPreforkChildVal)
	defer os.Unsetenv(envPreforkChildKey)

	// when
	result := IsChild()

	// then
	assert.True(t, result)
}
