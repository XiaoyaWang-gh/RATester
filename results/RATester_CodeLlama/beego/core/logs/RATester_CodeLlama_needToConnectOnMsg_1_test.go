package logs

import (
	"fmt"
	"io"
	"testing"
)

func TestNeedToConnectOnMsg_1(t *testing.T) {
	type fields struct {
		lg             *logWriter
		innerWriter    io.WriteCloser
		formatter      LogFormatter
		Formatter      string `json:"formatter"`
		ReconnectOnMsg bool   `json:"reconnectOnMsg"`
		Reconnect      bool   `json:"reconnect"`
		Net            string `json:"net"`
		Addr           string `json:"addr"`
		Level          int    `json:"level"`
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
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
			if got := c.needToConnectOnMsg(); got != tt.want {
				t.Errorf("needToConnectOnMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
