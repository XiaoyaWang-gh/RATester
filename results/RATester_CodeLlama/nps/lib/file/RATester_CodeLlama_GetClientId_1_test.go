package file

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetClientId_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	//given
	db := &JsonDb{}
	//when
	id := db.GetClientId()
	//then
	assert.Equal(t, int32(1), id)
}
