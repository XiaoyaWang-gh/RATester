package group

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
)

func TestUnRegister_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &HTTPGroup{
		group:           "test",
		groupKey:        "test",
		domain:          "test",
		location:        "test",
		routeByHTTPUser: "test",
		createFuncs:     make(map[string]vhost.CreateConnFunc),
		pxyNames:        []string{"test"},
		index:           0,
		ctl:             &HTTPGroupController{},
	}
	g.UnRegister("test")
	if len(g.createFuncs) != 0 {
		t.Fatal("UnRegister failed")
	}
}
