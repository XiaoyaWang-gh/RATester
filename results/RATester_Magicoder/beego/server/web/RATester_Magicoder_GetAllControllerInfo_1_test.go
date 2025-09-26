package web

import (
	"fmt"
	"testing"
)

func TestGetAllControllerInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrlRegister := &ControllerRegister{
		routers: map[string]*Tree{
			"test": &Tree{
				prefix: "test",
				leaves: []*leafInfo{
					{
						wildcards: []string{"id"},
						runObject: &ControllerInfo{
							pattern: "test/:id",
						},
					},
				},
			},
		},
	}

	routerInfos := ctrlRegister.GetAllControllerInfo()

	if len(routerInfos) != 1 {
		t.Errorf("Expected 1 router info, got %d", len(routerInfos))
	}

	if routerInfos[0].pattern != "test/:id" {
		t.Errorf("Expected pattern 'test/:id', got '%s'", routerInfos[0].pattern)
	}
}
