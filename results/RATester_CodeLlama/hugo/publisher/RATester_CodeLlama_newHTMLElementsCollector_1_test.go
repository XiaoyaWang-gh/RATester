package publisher

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestNewHTMLElementsCollector_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := config.BuildStats{}
	got := newHTMLElementsCollector(conf)
	if got == nil {
		t.Errorf("newHTMLElementsCollector() = %v, want %v", got, "not nil")
	}
}
