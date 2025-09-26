package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetClient_1(t *testing.T) {
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
		Id: 1,
	}
	db.JsonDb.Clients.Store(1, client)

	// 测试获取已存在的客户端
	got, err := db.GetClient(1)
	if err != nil {
		t.Errorf("GetClient() error = %v", err)
		return
	}
	if got.Id != 1 {
		t.Errorf("GetClient() got = %v, want %v", got, 1)
	}

	// 测试获取不存在的客户端
	_, err = db.GetClient(2)
	if err == nil {
		t.Errorf("GetClient() should return an error")
	}
}
