package logs

import (
	"fmt"
	"io"
	"testing"
)

func TestWriteMsg_6(t *testing.T) {
	type fields struct {
		lg             *logWriter
		innerWriter    io.WriteCloser
		formatter      LogFormatter
		Formatter      string
		ReconnectOnMsg bool
		Reconnect      bool
		Net            string
		Addr           string
		Level          int
	}
	type args struct {
		lm *LogMsg
	}
	tests := []struct {
		name    string
		fields  fields
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

			c := &connWriter{
				lg:             tt.fields.lg,
				innerWriter:    tt.fields.innerWriter,
				formatter:      tt.fields.formatter,
				Formatter:      tt.fields.Formatter,
				ReconnectOnMsg: tt.fields.ReconnectOnMsg,
				Reconnect:      tt.fields.Reconnect,
				Net:            tt.fields.Net,
				Addr:           tt.fields.Addr,
				Level:          tt.fields.Level,
			}
			if err := c.WriteMsg(tt.args.lm); (err != nil) != tt.wantErr {
				t.Errorf("connWriter.WriteMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
