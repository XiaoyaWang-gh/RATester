package alils

import (
	"fmt"
	"testing"
)

func TestGetLogs_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &LogStore{
		Name:           "test",
		TTL:            1,
		ShardCount:     1,
		CreateTime:     1,
		LastModifyTime: 1,
		project: &LogProject{
			Name:            "test",
			Endpoint:        "test",
			AccessKeyID:     "test",
			AccessKeySecret: "test",
		},
	}

	gl, nextCursor, err := store.GetLogs(1, "test", 1)
	if err != nil {
		t.Errorf("GetLogs failed, err: %v", err)
	}

	if gl == nil {
		t.Errorf("GetLogs failed, gl is nil")
	}

	if nextCursor == "" {
		t.Errorf("GetLogs failed, nextCursor is empty")
	}
}
