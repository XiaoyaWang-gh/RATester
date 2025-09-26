package schema

import (
	"fmt"
	"testing"
)

func TestSetAliasTag_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Encoder{
		cache: &cache{},
	}

	tag := "testTag"
	e.SetAliasTag(tag)

	if e.cache.tag != tag {
		t.Errorf("Expected tag to be %s, but got %s", tag, e.cache.tag)
	}
}
