package etcd

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestSaveConfigFile_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &EtcdConfiger{
		prefix: "test",
		client: &clientv3.Client{},
	}

	err := e.SaveConfigFile("test.json")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	expectedErr := errors.New("Unsupported operation")
	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}
}
