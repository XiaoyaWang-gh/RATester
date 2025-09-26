package file

import (
	"fmt"
	"testing"
)

func TestNewClient_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &DbUtils{
		JsonDb: &JsonDb{},
	}
	client := &Client{
		Id:          1,
		VerifyKey:   "test_key",
		WebUserName: "test_user",
	}
	err := db.NewClient(client)
	if err != nil {
		t.Errorf("TestNewClient failed, error: %v", err)
	}
}
