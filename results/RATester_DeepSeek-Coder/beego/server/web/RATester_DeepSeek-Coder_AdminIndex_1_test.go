package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestAdminIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &adminController{
		Controller: Controller{
			Ctx: &context.Context{},
		},
	}

	ctrl.AdminIndex()

	// Add your assertions here
}
