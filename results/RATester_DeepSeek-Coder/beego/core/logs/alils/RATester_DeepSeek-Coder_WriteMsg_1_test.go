package alils

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/logs"
)

func TestWriteMsg_1(t *testing.T) {
	type args struct {
		lm *logs.LogMsg
	}
	tests := []struct {
		name    string
		c       *aliLSWriter
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

			if err := tt.c.WriteMsg(tt.args.lm); (err != nil) != tt.wantErr {
				t.Errorf("aliLSWriter.WriteMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
