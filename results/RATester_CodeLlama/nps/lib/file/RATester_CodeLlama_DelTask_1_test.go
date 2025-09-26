package file

import (
	"fmt"
	"testing"
)

func TestDelTask_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var dbUtils DbUtils
	var id int = 1
	// act
	err := dbUtils.DelTask(id)
	// assert
	if err != nil {
		t.Errorf("DelTask() failed, err: %v", err)
	}
}
