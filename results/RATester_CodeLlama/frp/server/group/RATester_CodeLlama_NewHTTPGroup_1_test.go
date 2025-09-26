package group

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewHTTPGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &HTTPGroupController{
		groups: make(map[string]*HTTPGroup),
	}
	g := NewHTTPGroup(ctl)
	assert.NotNil(t, g)
}
