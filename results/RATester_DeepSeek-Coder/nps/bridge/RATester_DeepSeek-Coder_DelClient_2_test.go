package bridge

import (
	"fmt"
	"sync"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestDelClient_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Bridge{
		Client: sync.Map{},
	}

	// 添加一个客户端
	b.Client.Store(1, &Client{signal: &conn.Conn{}})

	// 删除客户端
	b.DelClient(1)

	// 检查客户端是否已删除
	_, ok := b.Client.Load(1)
	if ok {
		t.Errorf("Expected client to be deleted")
	}
}
