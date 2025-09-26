package file

import (
	"fmt"
	"sync"
	"testing"

	"github.com/zeebo/assert"
)

func TestUpdateTask_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var (
		dbUtils = &DbUtils{
			JsonDb: &JsonDb{
				Tasks: sync.Map{},
			},
		}
		tunnel = &Tunnel{
			Id: 1,
		}
	)

	// act
	err := dbUtils.UpdateTask(tunnel)

	// assert
	assert.NoError(t, err)
}
