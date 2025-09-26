package mem

import (
	"fmt"
	"testing"
)

func TestNewProxy_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &serverMetrics{}
	m.NewProxy("test", "test")
	if len(m.info.ProxyStatistics) != 1 {
		t.Errorf("NewProxy failed")
	}
}
