package client

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/client/proxy"
	"github.com/zeebo/assert"
)

func TestGetProxyStatus_1(t *testing.T) {
	t.Run("test case 1:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// arrange
		name := "test"
		status := &proxy.WorkingStatus{
			Name:       name,
			Type:       "test",
			Phase:      "test",
			Err:        "test",
			Cfg:        nil,
			RemoteAddr: "test",
		}
		getProxyStatusFunc := func(name string) (*proxy.WorkingStatus, bool) {
			return status, true
		}
		s := &statusExporterImpl{
			getProxyStatusFunc: getProxyStatusFunc,
		}
		// act
		result, ok := s.GetProxyStatus(name)
		// assert
		assert.Equal(t, status, result)
		assert.True(t, ok)
	})
}
