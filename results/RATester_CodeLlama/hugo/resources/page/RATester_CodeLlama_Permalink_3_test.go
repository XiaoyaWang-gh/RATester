package page

import (
	"fmt"
	"testing"
)

func TestPermalink_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := OutputFormat{
		permalink: "https://example.com/",
	}
	if got := o.Permalink(); got != o.permalink {
		t.Errorf("Permalink() = %v, want %v", got, o.permalink)
	}
}
