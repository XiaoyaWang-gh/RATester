package nathole

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestSendSidMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	conn, err := net.DialUDP("udp4", nil, &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080})
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	sid := "testSid"
	transactionID := "testTransactionID"
	addr := "127.0.0.1:8080"
	key := []byte("testKey")
	ttl := 10

	err = sendSidMessage(ctx, conn, sid, transactionID, addr, key, ttl)
	if err != nil {
		t.Fatal(err)
	}

	// Add assertions here to verify the expected behavior of the function
}
