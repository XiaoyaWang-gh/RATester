package server

import (
	"fmt"
	"sync"
	"testing"

	"ehang.io/nps/lib/file"
	"github.com/astaxie/beego"
)

func TestInitFromCsv_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	beego.AppConfig.Set("public_vkey", "test_vkey")
	file.GetDb().JsonDb.Tasks = sync.Map{}
	file.GetDb().JsonDb.Tasks.Store("test_id", &file.Tunnel{Status: true})

	InitFromCsv()

	_, ok := RunList.Load("test_id")
	if !ok {
		t.Errorf("Expected RunList to contain 'test_id'")
	}

	_, ok = file.GetDb().JsonDb.Tasks.Load("test_id")
	if !ok {
		t.Errorf("Expected Tasks to contain 'test_id'")
	}
}
