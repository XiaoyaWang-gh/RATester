package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/config/allconfig"
)

func TestCopyright_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	s.conf = &allconfig.Config{}
	s.conf.Copyright = "Copyright 2020"
	assert.Equal(t, "Copyright 2020", s.Copyright())
}
