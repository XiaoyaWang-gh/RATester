package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/config/allconfig"
)

func TestMainSections_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	s.conf = &allconfig.Config{}
	s.conf.C.MainSections = []string{"main"}
	assert.Equal(t, []string{"main"}, s.MainSections())
}
