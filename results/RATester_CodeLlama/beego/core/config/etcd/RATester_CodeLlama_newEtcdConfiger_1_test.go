package etcd

import (
	"fmt"
	"testing"
)

func TestNewEtcdConfiger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup your test
	// TODO: mock clientv3.Client
	// TODO: mock prefix
	// TODO: call newEtcdConfiger
	// TODO: verify your expectations
}
