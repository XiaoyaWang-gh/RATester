package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestWriteMsg_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	smtpWriter := &SMTPWriter{
		Username:           "testuser",
		Password:           "testpassword",
		Host:               "smtp.example.com:587",
		Subject:            "Test Subject",
		FromAddress:        "testuser@example.com",
		RecipientAddresses: []string{"recipient1@example.com", "recipient2@example.com"},
		Level:              1,
		InsecureSkipVerify: true,
	}

	logMsg := &LogMsg{
		Level: 1,
		Msg:   "Test Message",
		When:  time.Now(),
	}

	err := smtpWriter.WriteMsg(logMsg)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
