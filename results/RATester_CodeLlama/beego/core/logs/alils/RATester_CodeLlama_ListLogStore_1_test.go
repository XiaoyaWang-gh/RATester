package alils

import (
	"fmt"
	"testing"
)

func TestListLogStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &LogProject{
		Name:            "test",
		Endpoint:        "http://127.0.0.1:8888",
		AccessKeyID:     "accessKeyID",
		AccessKeySecret: "accessKeySecret",
	}

	storeNames, err := p.ListLogStore()
	if err != nil {
		t.Errorf("failed to list logstore: %v", err)
		return
	}

	if len(storeNames) == 0 {
		t.Errorf("failed to list logstore: %v", err)
		return
	}

	t.Logf("list logstore: %v", storeNames)
}
