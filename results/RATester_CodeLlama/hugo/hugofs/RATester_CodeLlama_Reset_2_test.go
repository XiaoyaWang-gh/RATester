package hugofs

import (
	"fmt"
	"testing"
)

func TestReset_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &createCountingFs{}
	c.onCreate("foo")
	c.Reset()
	if c.fileCount["foo"] != 0 {
		t.Errorf("expected file count to be 0, got %d", c.fileCount["foo"])
	}
}
