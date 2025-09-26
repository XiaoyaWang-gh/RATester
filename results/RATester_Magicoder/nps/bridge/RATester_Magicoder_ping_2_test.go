package bridge

import (
	"fmt"
	"testing"
	"time"

	"github.com/astaxie/beego/logs"
)

func Testping_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bridge := &Bridge{
		TunnelPort: 8080,
	}

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			arr := make([]int, 0)
			bridge.Client.Range(func(key, value interface{}) bool {
				v := value.(*Client)
				if v.tunnel == nil || v.signal == nil {
					v.retryTime += 1
					if v.retryTime >= 3 {
						arr = append(arr, key.(int))
					}
					return true
				}
				if v.tunnel.IsClose {
					arr = append(arr, key.(int))
				}
				return true
			})
			for _, v := range arr {
				logs.Info("the client %d closed", v)
				bridge.DelClient(v)
			}
		}
	}
}
