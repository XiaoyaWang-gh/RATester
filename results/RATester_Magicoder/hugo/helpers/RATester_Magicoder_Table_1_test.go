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

	stats := &ProcessingStats{
		Name:            "TestStats",
		Pages:           10,
		PaginatorPages:  5,
		Static:          3,
		ProcessedImages: 2,
		Files:           1,
		Aliases:         0,
		Cleaned:         0,
	}

	buf := &bytes.Buffer{}
	stats.Table(buf)

	expected := `+------------------+
| TestStats        |
+------------------+
| Pages            | 10 |
| PaginatorPages   | 5  |
| Static           | 3  |
| ProcessedImages | 2  |
| Files            | 1  |
| Aliases          | 0  |
| Cleaned          | 0  |
+------------------+
`

	if buf.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, buf.String())
	}
}
