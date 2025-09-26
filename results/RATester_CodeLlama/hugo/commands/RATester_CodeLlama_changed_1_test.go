package commands

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestChanged_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &fileChangeDetector{}
	f.current = map[string]uint64{"a": 1, "b": 2, "c": 3}
	f.prev = map[string]uint64{"a": 1, "b": 2, "c": 3}
	f.irrelevantRe = regexp.MustCompile("^a$")
	if got := f.changed(); !reflect.DeepEqual(got, []string{"a"}) {
		t.Errorf("changed() = %v, want %v", got, []string{"a"})
	}
}
