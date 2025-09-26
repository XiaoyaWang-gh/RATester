package alils

import (
	"fmt"
	"sync"
	"testing"

	"github.com/beego/beego/v2/core/logs"
)

func TestDestroy_1(t *testing.T) {
	type fields struct {
		store    *LogStore
		group    []*LogGroup
		withMap  bool
		groupMap map[string]*LogGroup
		lock     *sync.Mutex
		Config
		formatter logs.LogFormatter
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &aliLSWriter{
				store:     tt.fields.store,
				group:     tt.fields.group,
				withMap:   tt.fields.withMap,
				groupMap:  tt.fields.groupMap,
				lock:      tt.fields.lock,
				Config:    tt.fields.Config,
				formatter: tt.fields.formatter,
			}
			c.Destroy()
		})
	}
}
