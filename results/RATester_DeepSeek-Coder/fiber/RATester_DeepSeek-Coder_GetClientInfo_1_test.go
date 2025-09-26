package fiber

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"
)

func TestGetClientInfo_1(t *testing.T) {
	type args struct {
		info *tls.ClientHelloInfo
	}
	tests := []struct {
		name    string
		t       *TLSHandler
		args    args
		want    *tls.Certificate
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

			got, err := tt.t.GetClientInfo(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("TLSHandler.GetClientInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TLSHandler.GetClientInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
