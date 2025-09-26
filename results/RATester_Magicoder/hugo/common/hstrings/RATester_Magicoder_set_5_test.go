package hstrings

import (
	"fmt"
	"regexp"
	"testing"
)

func Testset_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rc := &regexpCache{
		re: make(map[string]*regexp.Regexp),
	}

	key := "testKey"
	re, err := regexp.Compile(`testRegexp`)
	if err != nil {
		t.Fatalf("Failed to compile regexp: %v", err)
	}

	rc.set(key, re)

	if _, ok := rc.re[key]; !ok {
		t.Errorf("Expected key %s to be in the cache", key)
	}
}
