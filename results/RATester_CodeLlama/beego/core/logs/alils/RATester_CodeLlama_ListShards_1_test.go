package alils

import (
	"fmt"
	"testing"
)

func TestListShards_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &LogStore{
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
	shardIDs, err := s.ListShards()
	if err != nil {
		t.Errorf("ListShards failed, err: %v", err)
	}
	if len(shardIDs) != 1 {
		t.Errorf("ListShards failed, shardIDs: %v", shardIDs)
	}
}
