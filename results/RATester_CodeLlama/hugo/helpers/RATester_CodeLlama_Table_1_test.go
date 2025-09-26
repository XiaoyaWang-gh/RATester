package helpers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &ProcessingStats{
		Name: "test",
	}

	var buf bytes.Buffer
	s.Table(&buf)

	expected := `
+-------+-------+
|       | test  |
+-------+-------+
| Pages | 0     |
|       |       |
+-------+-------+
`

	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}
