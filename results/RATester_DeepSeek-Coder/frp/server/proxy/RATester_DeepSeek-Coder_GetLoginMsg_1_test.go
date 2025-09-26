package proxy

import (
	"fmt"
	"testing"
	"time"

	"github.com/fatedier/frp/pkg/msg"
)

func TestGetLoginMsg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pxy := &BaseProxy{
		name: "test",
		loginMsg: &msg.Login{
			Version:      "v1",
			Hostname:     "localhost",
			Os:           "linux",
			Arch:         "amd64",
			User:         "test",
			PrivilegeKey: "privilege_key",
			Timestamp:    time.Now().Unix(),
			RunID:        "run_id",
			Metas: map[string]string{
				"meta1": "value1",
				"meta2": "value2",
			},
			PoolCount: 10,
		},
	}

	loginMsg := pxy.GetLoginMsg()

	if loginMsg.Version != "v1" {
		t.Errorf("Expected version 'v1', got '%s'", loginMsg.Version)
	}

	if loginMsg.Hostname != "localhost" {
		t.Errorf("Expected hostname 'localhost', got '%s'", loginMsg.Hostname)
	}

	if loginMsg.Os != "linux" {
		t.Errorf("Expected os 'linux', got '%s'", loginMsg.Os)
	}

	if loginMsg.Arch != "amd64" {
		t.Errorf("Expected arch 'amd64', got '%s'", loginMsg.Arch)
	}

	if loginMsg.User != "test" {
		t.Errorf("Expected user 'test', got '%s'", loginMsg.User)
	}

	if loginMsg.PrivilegeKey != "privilege_key" {
		t.Errorf("Expected privilege key 'privilege_key', got '%s'", loginMsg.PrivilegeKey)
	}

	if loginMsg.PoolCount != 10 {
		t.Errorf("Expected pool count 10, got '%d'", loginMsg.PoolCount)
	}
}
