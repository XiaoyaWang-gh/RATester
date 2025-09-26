package alils

import (
	"fmt"
	"testing"
)

func TestGetLogStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	project := &LogProject{
		Name:            "test_project",
		Endpoint:        "http://localhost:8080",
		AccessKeyID:     "test_key_id",
		AccessKeySecret: "test_key_secret",
	}

	storeName := "test_store"
	_, err := project.GetLogStore(storeName)
	if err != nil {
		t.Errorf("GetLogStore failed: %v", err)
	}
}
