package proxy

import (
	"fmt"
	"sync"
	"testing"
)

func TestDel_3(t *testing.T) {
	type fields struct {
		pxys map[string]Proxy
		mu   sync.RWMutex
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			pm := &Manager{
				pxys: tt.fields.pxys,
				mu:   tt.fields.mu,
			}
			pm.Del(tt.args.name)
		})
	}
}
