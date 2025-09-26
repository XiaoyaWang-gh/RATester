package task

import (
	"fmt"
	"testing"
)

func TestExecute_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &runTaskCommand{}
	params := []interface{}{"task1"}
	result := r.Execute(params...)
	if result.IsSuccess() {
		t.Errorf("Expected failure, but got success")
	}
}
