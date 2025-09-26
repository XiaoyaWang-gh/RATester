package publisher

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestnewHTMLElementsCollector_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := config.BuildStats{
		Enable:         true,
		DisableTags:    false,
		DisableClasses: false,
		DisableIDs:     false,
	}

	collector := newHTMLElementsCollector(conf)

	if collector == nil {
		t.Error("Expected collector to be not nil")
	}

	if collector.conf.Enable != conf.Enable {
		t.Errorf("Expected collector.conf.Enable to be %v, but got %v", conf.Enable, collector.conf.Enable)
	}

	if collector.conf.DisableTags != conf.DisableTags {
		t.Errorf("Expected collector.conf.DisableTags to be %v, but got %v", conf.DisableTags, collector.conf.DisableTags)
	}

	if collector.conf.DisableClasses != conf.DisableClasses {
		t.Errorf("Expected collector.conf.DisableClasses to be %v, but got %v", conf.DisableClasses, collector.conf.DisableClasses)
	}

	if collector.conf.DisableIDs != conf.DisableIDs {
		t.Errorf("Expected collector.conf.DisableIDs to be %v, but got %v", conf.DisableIDs, collector.conf.DisableIDs)
	}

	if len(collector.elementSet) != 0 {
		t.Errorf("Expected collector.elementSet to be empty, but got %v", collector.elementSet)
	}
}
