package logs

import (
	"fmt"
	"testing"
)

func TestWriteMsg_10(t *testing.T) {
	type args struct {
		lm *LogMsg
	}
	tests := []struct {
		name    string
		bl      *BeeLogger
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

			if err := tt.bl.writeMsg(tt.args.lm); (err != nil) != tt.wantErr {
				t.Errorf("BeeLogger.writeMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
