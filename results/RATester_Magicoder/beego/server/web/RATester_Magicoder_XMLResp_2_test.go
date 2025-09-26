package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestXMLResp_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
	}

	data := struct {
		Name string `xml:"name"`
	}{
		Name: "test",
	}

	err := ctrl.XMLResp(data)
	if err != nil {
		t.Errorf("XMLResp() error = %v", err)
	}
}
