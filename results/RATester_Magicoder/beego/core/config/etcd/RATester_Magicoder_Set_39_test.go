package etcd

import (
	"fmt"
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestSet_39(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	etcdConfiger := &EtcdConfiger{
		prefix: "test_prefix",
		client: &clientv3.Client{},
	}

	err := etcdConfiger.Set("test_key", "test_value")
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if err.Error() != "Unsupported operation" {
		t.Errorf("Expected error message 'Unsupported operation', but got '%s'", err.Error())
	}
}
