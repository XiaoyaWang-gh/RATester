package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestFlush_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	st := &SessionStore{}
	st.Flush(context.Background())
	if len(st.values) != 0 {
		t.Errorf("Flush() failed")
	}
}
