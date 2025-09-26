package template

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTJSTmpl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c context
	var s []byte
	var k int
	for {
		i := k + bytes.IndexAny(s[k:], "`\\$")
		if i < k {
			break
		}
		switch s[i] {
		case '\\':
			i++
			if i == len(s) {
				if c.state != stateError {
					t.Errorf("unfinished escape sequence in JS string: %q", s)
				}
				return
			}
		case '$':
			if len(s) >= i+2 && s[i+1] == '{' {
				c.jsBraceDepth = append(c.jsBraceDepth, 0)
				c.state = stateJS
				return
			}
		case '`':
			c.state = stateJS
			return
		}
		k = i + 1
	}

	if c.state != stateJS {
		t.Errorf("unfinished JS string: %q", s)
	}
}
