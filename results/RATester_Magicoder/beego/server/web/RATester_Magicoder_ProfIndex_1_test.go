package web

import (
	"fmt"
	"testing"
)

func TestProfIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &adminController{
		servers: []*HttpServer{},
	}

	ctrl.ProfIndex()

	// Add assertions here
}
