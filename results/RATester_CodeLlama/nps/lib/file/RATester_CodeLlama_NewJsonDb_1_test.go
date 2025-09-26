package file

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewJsonDb_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	runPath := "./testdata"
	// act
	db := NewJsonDb(runPath)
	// assert
	assert.NotNil(t, db)
	assert.Equal(t, runPath, db.RunPath)
	assert.Equal(t, filepath.Join(runPath, "conf", "tasks.json"), db.TaskFilePath)
	assert.Equal(t, filepath.Join(runPath, "conf", "hosts.json"), db.HostFilePath)
	assert.Equal(t, filepath.Join(runPath, "conf", "clients.json"), db.ClientFilePath)
}
