package logs

import (
	"fmt"
	"testing"
)

func TestsendMail_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	writer := &SMTPWriter{
		Username:           "testuser",
		Password:           "testpassword",
		Host:               "testhost",
		Subject:            "testsubject",
		FromAddress:        "testfromaddress",
		RecipientAddresses: []string{"testrecipient1", "testrecipient2"},
		Level:              1,
		InsecureSkipVerify: true,
	}

	hostAddressWithPort := "testhost:25"
	auth := writer.getSMTPAuth(hostAddressWithPort)
	msgContent := []byte("test message content")

	err := writer.sendMail(hostAddressWithPort, auth, writer.FromAddress, writer.RecipientAddresses, msgContent)
	if err != nil {
		t.Errorf("sendMail failed: %v", err)
	}
}
