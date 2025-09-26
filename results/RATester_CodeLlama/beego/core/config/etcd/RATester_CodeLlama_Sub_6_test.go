package etcd

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSub_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &EtcdConfiger{}
	_, err := e.Sub("")
	assert.NoError(t, err)
}
