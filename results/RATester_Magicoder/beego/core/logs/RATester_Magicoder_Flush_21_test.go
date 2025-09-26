package logs

import (
	"fmt"
	"testing"
)

func TestFlush_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	writer := &JLWriter{
		AuthorName: "test",
		Title:      "test",
		WebhookURL: "http://test.com",
	}
	writer.Flush()
}
