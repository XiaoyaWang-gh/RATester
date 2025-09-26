package bridge

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TesttypeDeal_1(t *testing.T) {
	type args struct {
		typeVal string
		c       *conn.Conn
		id      int
		vs      string
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
			s.typeDeal(tt.args.typeVal, tt.args.c, tt.args.id, tt.args.vs)
		})
	}
}
