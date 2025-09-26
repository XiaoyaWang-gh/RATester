package etcd

import (
	"fmt"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestSub_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	etcdConfiger := &EtcdConfiger{
		prefix: "test/",
		client: client,
	}

	subConfiger, err := etcdConfiger.Sub("subkey")
	if err != nil {
		t.Fatal(err)
	}

	if subConfiger == nil {
		t.Fatal("SubConfiger should not be nil")
	}

	// Add more specific tests here
}
