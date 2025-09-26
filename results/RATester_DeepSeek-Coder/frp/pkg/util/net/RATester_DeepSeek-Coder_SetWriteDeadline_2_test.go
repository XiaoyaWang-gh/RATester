package net

import (
	"fmt"
	"testing"
	"time"
)

func TestSetWriteDeadline_2(t *testing.T) {
	type args struct {
		deadline time.Time
	}
	tests := []struct {
		name    string
		c       *FakeUDPConn
		args    args
		wantErr bool
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

			if err := tt.c.SetWriteDeadline(tt.args.deadline); (err != nil) != tt.wantErr {
				t.Errorf("FakeUDPConn.SetWriteDeadline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
