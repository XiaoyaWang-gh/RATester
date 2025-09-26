package logs

import (
	"fmt"
	"net/smtp"
	"reflect"
	"testing"
)

func TestGetSMTPAuth_1(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		s    *SMTPWriter
		args args
		want smtp.Auth
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

			if got := tt.s.getSMTPAuth(tt.args.host); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SMTPWriter.getSMTPAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
