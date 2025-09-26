package logs

import (
	"fmt"
	"testing"
)

func TestFlush_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	writer := &SLACKWriter{
		WebhookURL: "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX",
		Level:      1,
		Formatter:  "json",
	}

	writer.Flush()
}
