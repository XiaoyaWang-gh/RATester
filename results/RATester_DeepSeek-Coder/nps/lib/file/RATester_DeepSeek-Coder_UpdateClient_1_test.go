package file

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/rate"
)

func TestUpdateClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &DbUtils{
		JsonDb: &JsonDb{},
	}

	client := &Client{
		Id:        1,
		Addr:      "127.0.0.1",
		Rate:      rate.NewRate(1024),
		RateLimit: 1024,
	}

	err := db.UpdateClient(client)
	if err != nil {
		t.Errorf("UpdateClient failed, err: %v", err)
	}

	storedClient, err := db.GetClient(1)
	if err != nil {
		t.Errorf("GetClient failed, err: %v", err)
	}

	if storedClient.Id != client.Id {
		t.Errorf("Expected Id %d, got %d", client.Id, storedClient.Id)
	}

	if storedClient.Addr != client.Addr {
		t.Errorf("Expected Addr %s, got %s", client.Addr, storedClient.Addr)
	}

	if storedClient.RateLimit != client.RateLimit {
		t.Errorf("Expected RateLimit %d, got %d", client.RateLimit, storedClient.RateLimit)
	}
}
