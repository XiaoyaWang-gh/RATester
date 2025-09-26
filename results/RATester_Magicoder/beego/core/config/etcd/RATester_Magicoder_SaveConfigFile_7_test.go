package etcd

import (
	"fmt"
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestSaveConfigFile_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	etcdConfiger := &EtcdConfiger{
		prefix: "testPrefix",
		client: &clientv3.Client{},
	}

	err := etcdConfiger.SaveConfigFile("testFile")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
