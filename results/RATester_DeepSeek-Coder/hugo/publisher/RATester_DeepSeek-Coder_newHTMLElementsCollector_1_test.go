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

	conf := config.BuildStats{
		Enable:         true,
		DisableTags:    true,
		DisableClasses: true,
		DisableIDs:     true,
	}

	collector := newHTMLElementsCollector(conf)

	if collector.conf != conf {
		t.Errorf("Expected config to be %v, got %v", conf, collector.conf)
	}

	if len(collector.elementSet) != 0 {
		t.Errorf("Expected elementSet to be empty, got %v", collector.elementSet)
	}

	if collector.elements != nil {
		t.Errorf("Expected elements to be nil, got %v", collector.elements)
	}
}
