package maps

import (
	"fmt"
	"testing"
)

func TestForEeach_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache[string, string]{
		m: map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	}
	f := func(k string, v string) {
		t.Logf("key: %s, value: %s", k, v)
	}
	c.ForEeach(f)
}
