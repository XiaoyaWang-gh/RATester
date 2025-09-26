package ledis

import (
	"fmt"
	"testing"
)

func TestSessionRelease_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ls := &SessionStore{}
	ls.sid = "123456"
	ls.values = make(map[interface{}]interface{})
	ls.values["name"] = "astaxie"
	ls.maxlifetime = 10
	ls.SessionRelease(nil, nil)
	if len(ls.values) != 0 {
		t.Error("SessionRelease failed")
	}
}
