package logs

import (
	"fmt"
	"testing"
)

func TestWriteMsg_8(t *testing.T) {
	type args struct {
		lm *LogMsg
	}
	tests := []struct {
		name    string
		s       *JLWriter
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

			if err := tt.s.WriteMsg(tt.args.lm); (err != nil) != tt.wantErr {
				t.Errorf("JLWriter.WriteMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
