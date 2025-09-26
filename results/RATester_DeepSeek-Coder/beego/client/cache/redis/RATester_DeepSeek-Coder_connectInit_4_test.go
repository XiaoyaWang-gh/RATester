package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

func TestConnectInit_4(t *testing.T) {
	type fields struct {
		p        *redis.Pool
		conninfo string
		dbNum    int
		key      string
		password string
		maxIdle  int
		timeout  time.Duration
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

			rc := &Cache{
				p:        tt.fields.p,
				conninfo: tt.fields.conninfo,
				dbNum:    tt.fields.dbNum,
				key:      tt.fields.key,
				password: tt.fields.password,
				maxIdle:  tt.fields.maxIdle,
				timeout:  tt.fields.timeout,
			}
			rc.connectInit()
		})
	}
}
