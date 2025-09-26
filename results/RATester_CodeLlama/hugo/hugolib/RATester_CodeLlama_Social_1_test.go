package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/config/allconfig"
)

func TestSocial_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	s.conf = &allconfig.Config{}
	s.conf.Social = map[string]string{"twitter": "https://twitter.com/gohugoio"}
	assert.Equal(t, s.Social(), s.conf.Social)
}
