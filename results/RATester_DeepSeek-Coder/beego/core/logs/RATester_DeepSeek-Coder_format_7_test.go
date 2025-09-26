package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat_7(t *testing.T) {
	type fields struct {
		RemoteAddr     string
		RequestTime    time.Time
		RequestMethod  string
		Request        string
		ServerProtocol string
		Host           string
		Status         int
		BodyBytesSent  int64
		ElapsedTime    time.Duration
		HTTPReferrer   string
		HTTPUserAgent  string
		RemoteUser     string
	}
	type args struct {
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
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

			r := &AccessLogRecord{
				RemoteAddr:     tt.fields.RemoteAddr,
				RequestTime:    tt.fields.RequestTime,
				RequestMethod:  tt.fields.RequestMethod,
				Request:        tt.fields.Request,
				ServerProtocol: tt.fields.ServerProtocol,
				Host:           tt.fields.Host,
				Status:         tt.fields.Status,
				BodyBytesSent:  tt.fields.BodyBytesSent,
				ElapsedTime:    tt.fields.ElapsedTime,
				HTTPReferrer:   tt.fields.HTTPReferrer,
				HTTPUserAgent:  tt.fields.HTTPUserAgent,
				RemoteUser:     tt.fields.RemoteUser,
			}
			if got := r.format(tt.args.format); got != tt.want {
				t.Errorf("AccessLogRecord.format() = %v, want %v", got, tt.want)
			}
		})
	}
}
