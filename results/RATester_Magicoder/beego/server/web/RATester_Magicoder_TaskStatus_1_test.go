package web

import (
	"fmt"
	"testing"
)

func TestTaskStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &adminController{
		servers: []*HttpServer{},
	}

	ctrl.TaskStatus()

	// Add assertions here
}
