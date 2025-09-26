package forwardedheaders

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewXForwarded_1(t *testing.T) {
	type args struct {
		insecure   bool
		trustedIps []string
		next       http.Handler
	}
	tests := []struct {
		name    string
		args    args
		want    *XForwarded
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

			got, err := NewXForwarded(tt.args.insecure, tt.args.trustedIps, tt.args.next)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewXForwarded() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewXForwarded() = %v, want %v", got, tt.want)
			}
		})
	}
}
