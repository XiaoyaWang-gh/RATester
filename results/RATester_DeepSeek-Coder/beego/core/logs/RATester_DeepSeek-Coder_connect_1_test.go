package logs

import (
	"fmt"
	"io"
	"testing"
)

func TestConnect_1(t *testing.T) {
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
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				Net:  "tcp",
				Addr: "localhost:8080",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			fields: fields{
				Net:  "tcp",
				Addr: "localhost:8081",
			},
			wantErr: true,
		},
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
			if err := c.connect(); (err != nil) != tt.wantErr {
				t.Errorf("connWriter.connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
