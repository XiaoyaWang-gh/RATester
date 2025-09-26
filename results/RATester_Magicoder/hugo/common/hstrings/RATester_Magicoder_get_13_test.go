package hstrings

import (
	"fmt"
	"regexp"
	"testing"
)

func Testget_13(t *testing.T) {
	rc := &regexpCache{
		re: map[string]*regexp.Regexp{
			"key1": regexp.MustCompile("value1"),
			"key2": regexp.MustCompile("value2"),
		},
	}

	tests := []struct {
		name     string
		key      string
		expected *regexp.Regexp
		ok       bool
	}{
		{
			name:     "Existing key",
			key:      "key1",
			expected: regexp.MustCompile("value1"),
			ok:       true,
		},
		{
			name:     "Non-existing key",
			key:      "key3",
			expected: nil,
			ok:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			re, ok := rc.get(test.key)
			if ok != test.ok {
				t.Errorf("Expected ok to be %v, but got %v", test.ok, ok)
			}
			if re != test.expected {
				t.Errorf("Expected regexp to be %v, but got %v", test.expected, re)
			}
		})
	}
}
