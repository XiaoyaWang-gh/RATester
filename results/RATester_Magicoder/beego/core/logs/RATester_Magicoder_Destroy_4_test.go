package logs

import (
	"fmt"
	"testing"
)

func TestDestroy_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	writer := &JLWriter{
		AuthorName: "test",
		Title:      "test",
	}
	writer.Destroy()
}
