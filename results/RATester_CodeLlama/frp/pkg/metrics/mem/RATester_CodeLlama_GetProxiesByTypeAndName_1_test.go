package mem

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestGetProxiesByTypeAndName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &serverMetrics{
		info: &ServerStatistics{
			ProxyStatistics: map[string]*ProxyStatistics{
				"proxy1": {
					ProxyType: "http",
				},
				"proxy2": {
					ProxyType: "tcp",
				},
			},
		},
	}
	res := m.GetProxiesByTypeAndName("http", "proxy1")
	assert.NotNil(t, res)
	assert.Equal(t, "proxy1", res.Name)
	assert.Equal(t, "http", res.Type)
	assert.Equal(t, int64(0), res.TodayTrafficIn)
	assert.Equal(t, int64(0), res.TodayTrafficOut)
	assert.Equal(t, int64(0), res.CurConns)
	assert.Equal(t, "", res.LastStartTime)
	assert.Equal(t, "", res.LastCloseTime)

	res = m.GetProxiesByTypeAndName("tcp", "proxy2")
	assert.NotNil(t, res)
	assert.Equal(t, "proxy2", res.Name)
	assert.Equal(t, "tcp", res.Type)
	assert.Equal(t, int64(0), res.TodayTrafficIn)
	assert.Equal(t, int64(0), res.TodayTrafficOut)
	assert.Equal(t, int64(0), res.CurConns)
	assert.Equal(t, "", res.LastStartTime)
	assert.Equal(t, "", res.LastCloseTime)

	res = m.GetProxiesByTypeAndName("http", "proxy2")
	assert.Nil(t, res)
}
