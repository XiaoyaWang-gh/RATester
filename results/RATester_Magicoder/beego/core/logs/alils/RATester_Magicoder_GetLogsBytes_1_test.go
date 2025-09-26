package alils

import (
	"fmt"
	"testing"
)

func TestGetLogsBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &LogStore{
		Name: "test",
		project: &LogProject{
			Name:            "test",
			Endpoint:        "http://localhost:8080",
			AccessKeyID:     "test",
			AccessKeySecret: "test",
		},
	}

	_, _, err := store.GetLogsBytes(1, "test", 10)
	if err != nil {
		t.Errorf("GetLogsBytes failed: %v", err)
	}
}
