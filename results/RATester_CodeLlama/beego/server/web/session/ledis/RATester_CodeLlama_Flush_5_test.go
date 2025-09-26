package ledis

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestFlush_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ls := &SessionStore{}
	ls.lock = sync.RWMutex{}
	ls.values = make(map[interface{}]interface{})
	ls.values["test"] = "test"
	ls.Flush(context.Background())
	if len(ls.values) != 0 {
		t.Error("Flush failed")
	}
}
