package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetClientIdByVkey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &DbUtils{
		JsonDb: &JsonDb{
			Clients: sync.Map{},
		},
	}

	// 添加一个客户端
	client := &Client{
		Id:        1,
		VerifyKey: "test_vkey",
	}
	db.JsonDb.Clients.Store(1, client)

	// 测试获取客户端ID
	id, err := db.GetClientIdByVkey("test_vkey")
	if err != nil {
		t.Errorf("GetClientIdByVkey() error = %v", err)
		return
	}
	if id != 1 {
		t.Errorf("GetClientIdByVkey() = %v, want %v", id, 1)
	}

	// 测试未找到客户端
	id, err = db.GetClientIdByVkey("not_exist")
	if err == nil {
		t.Errorf("GetClientIdByVkey() error = %v", err)
		return
	}
	if id != 0 {
		t.Errorf("GetClientIdByVkey() = %v, want %v", id, 0)
	}
}
