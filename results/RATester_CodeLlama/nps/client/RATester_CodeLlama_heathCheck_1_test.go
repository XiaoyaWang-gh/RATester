package client

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
	"ehang.io/nps/lib/file"
)

func TestHeathCheck_1(t *testing.T) {
	type args struct {
		healths []*file.Health
		c       *conn.Conn
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := heathCheck(tt.args.healths, tt.args.c); got != tt.want {
				t.Errorf("heathCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
