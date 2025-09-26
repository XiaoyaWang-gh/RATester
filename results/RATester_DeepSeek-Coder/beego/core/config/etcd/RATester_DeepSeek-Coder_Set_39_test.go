package etcd

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestSet_39(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatalf("Failed to create etcd client: %v", err)
	}
	defer etcdClient.Close()

	etcdConfiger := &EtcdConfiger{
		prefix: "test",
		client: etcdClient,
	}

	key := "key1"
	val := "value1"

	err = etcdConfiger.Set(key, val)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	expectedErr := errors.New("Unsupported operation")
	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}
}
