package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetClient_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &JsonDb{
		Clients: sync.Map{},
	}

	// 测试不存在的客户端
	_, err := db.GetClient(1)
	if err == nil {
		t.Errorf("期望得到错误，但没有得到")
	}

	// 添加一个客户端
	client := &Client{Id: 1}
	db.Clients.Store(1, client)

	// 测试存在的客户端
	c, err := db.GetClient(1)
	if err != nil {
		t.Errorf("期望得到客户端，但得到错误: %v", err)
	}
	if c.Id != 1 {
		t.Errorf("期望得到客户端ID为1，但得到ID为%d", c.Id)
	}
}
