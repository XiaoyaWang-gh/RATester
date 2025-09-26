package requestdecorator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/miekg/dns"
)

func TestGetRecord_1(t *testing.T) {
	type args struct {
		client *dns.Client
		msg    *dns.Msg
		server string
		port   string
	}
	tests := []struct {
		name    string
		args    args
		want    *cnameResolv
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

			got, err := getRecord(tt.args.client, tt.args.msg, tt.args.server, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
