package process

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNew_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := New("test", []string{"test"})
	assert.NotNil(t, p)
}
