package alils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListShards_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"shardID": 1}, {"shardID": 2}]`))
	}))
	defer server.Close()

	s := &LogStore{
		Name: "test",
		project: &LogProject{
			Endpoint: server.URL,
		},
	}

	shardIDs, err := s.ListShards()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(shardIDs) != 2 {
		t.Errorf("expected 2 shards, got %d", len(shardIDs))
	}

	if shardIDs[0] != 1 || shardIDs[1] != 2 {
		t.Errorf("unexpected shard IDs: %v", shardIDs)
	}
}
