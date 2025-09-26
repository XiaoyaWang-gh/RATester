package logs

import (
	"fmt"
	"testing"
)

func TestFormat_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	writer := &SMTPWriter{
		Username: "testuser",
		Password: "testpass",
		Host:     "testhost",
	}
	msg := &LogMsg{
		Msg: "test message",
	}
	formattedMsg := writer.Format(msg)
	if formattedMsg != msg.OldStyleFormat() {
		t.Errorf("Expected %s, got %s", msg.OldStyleFormat(), formattedMsg)
	}
}
