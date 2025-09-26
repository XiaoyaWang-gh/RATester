package httpcache

import (
	"fmt"
	"testing"
	"time"
)

func TestMarshalJSON_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := PollConfig{
		For:     GlobMatcher{Includes: []string{"foo"}},
		Disable: true,
		Low:     1 * time.Second,
		High:    10 * time.Second,
	}
	b, err := c.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"For":{"Includes":["foo"]},"Disable":true,"Low":"1s","High":"10s"}` {
		t.Errorf("unexpected JSON: %s", b)
	}
}
