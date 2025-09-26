package orm

import (
	"fmt"
	"testing"
)

func Testadd_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ac := &_dbCache{
		cache: make(map[string]*alias),
	}

	al := &alias{
		Name: "test",
	}

	added := ac.add("test", al)
	if !added {
		t.Error("Expected true, got false")
	}

	added = ac.add("test", al)
	if added {
		t.Error("Expected false, got true")
	}
}
