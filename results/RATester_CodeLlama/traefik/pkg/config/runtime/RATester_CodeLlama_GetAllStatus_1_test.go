package runtime

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetAllStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &ServiceInfo{}
	s.serverStatusMu.Lock()
	s.serverStatus = map[string]string{
		"server1": "status1",
		"server2": "status2",
	}
	s.serverStatusMu.Unlock()

	allStatus := s.GetAllStatus()
	assert.Equal(t, map[string]string{
		"server1": "status1",
		"server2": "status2",
	}, allStatus)
}
