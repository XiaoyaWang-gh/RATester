package bridge

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
	"ehang.io/nps/lib/file"
)

func TestgetConfig_1(t *testing.T) {
	type args struct {
		c      *conn.Conn
		isPub  bool
		client *file.Client
	}
	tests := []struct {
		name string
		args args
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

			s := &Bridge{}
			s.getConfig(tt.args.c, tt.args.isPub, tt.args.client)
		})
	}
}
