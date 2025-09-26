package hstrings

import (
	"fmt"
	"regexp"
	"testing"
)

func TestSet_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rc := &regexpCache{}
	rc.mu.Lock()
	rc.re = make(map[string]*regexp.Regexp)
	rc.mu.Unlock()
	key := "key"
	re := &regexp.Regexp{}
	rc.set(key, re)
	rc.mu.RLock()
	if rc.re[key] != re {
		t.Errorf("rc.re[key] = %v, want %v", rc.re[key], re)
	}
	rc.mu.RUnlock()
}
