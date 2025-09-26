package logs

import (
	"fmt"
	"testing"
)

func TestWriteMsg_4(t *testing.T) {
	type args struct {
		lm *LogMsg
	}
	tests := []struct {
		name    string
		f       *multiFileLogWriter
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

			if err := tt.f.WriteMsg(tt.args.lm); (err != nil) != tt.wantErr {
				t.Errorf("multiFileLogWriter.WriteMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
