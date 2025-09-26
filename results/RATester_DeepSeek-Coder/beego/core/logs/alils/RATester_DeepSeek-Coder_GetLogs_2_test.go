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
		TTL:            10,
		ShardCount:     1,
		CreateTime:     1514764800,
		LastModifyTime: 1514764800,
	}

	shardID := 0
	cursor := ""
	logGroupMaxCount := 10

	gl, nextCursor, err := store.GetLogs(shardID, cursor, logGroupMaxCount)
	if err != nil {
		t.Errorf("GetLogs failed, err: %v", err)
		return
	}

	if gl == nil {
		t.Errorf("GetLogs failed, log group list is nil")
		return
	}

	if nextCursor == "" {
		t.Errorf("GetLogs failed, next cursor is empty")
		return
	}

	t.Logf("GetLogs success, log group list: %v, next cursor: %s", gl, nextCursor)
}
