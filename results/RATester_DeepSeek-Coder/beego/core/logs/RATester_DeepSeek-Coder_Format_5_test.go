package logs

import (
	"fmt"
	"testing"
)

func TestFormat_5(t *testing.T) {
	type fields struct {
		Username           string
		Password           string
		Host               string
		Subject            string
		FromAddress        string
		RecipientAddresses []string
		Level              int
		InsecureSkipVerify bool
		formatter          LogFormatter
		Formatter          string
	}
	type args struct {
		lm *LogMsg
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

			s := &SMTPWriter{
				Username:           tt.fields.Username,
				Password:           tt.fields.Password,
				Host:               tt.fields.Host,
				Subject:            tt.fields.Subject,
				FromAddress:        tt.fields.FromAddress,
				RecipientAddresses: tt.fields.RecipientAddresses,
				Level:              tt.fields.Level,
				InsecureSkipVerify: tt.fields.InsecureSkipVerify,
				formatter:          tt.fields.formatter,
				Formatter:          tt.fields.Formatter,
			}
			if got := s.Format(tt.args.lm); got != tt.want {
				t.Errorf("SMTPWriter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
