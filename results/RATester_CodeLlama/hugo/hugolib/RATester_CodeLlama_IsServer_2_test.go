package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config/allconfig"
)

func TestIsServer_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	s.conf = &allconfig.Config{}
	s.conf.Internal.Running = true
	if !s.IsServer() {
		t.Errorf("IsServer() = %v, want %v", false, true)
	}
}
