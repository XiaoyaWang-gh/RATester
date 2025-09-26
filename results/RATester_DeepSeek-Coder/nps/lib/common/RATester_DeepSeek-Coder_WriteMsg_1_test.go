package common

import (
	"fmt"
	"testing"
	"time"
)

func TestWriteMsg_1(t *testing.T) {
	type args struct {
		when  time.Time
		msg   string
		level int
	}
	tests := []struct {
		name    string
		lg      *StoreMsg
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

			if err := tt.lg.WriteMsg(tt.args.when, tt.args.msg, tt.args.level); (err != nil) != tt.wantErr {
				t.Errorf("StoreMsg.WriteMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
