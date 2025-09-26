package headers

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewHeader_1(t *testing.T) {
	type args struct {
		next http.Handler
		cfg  dynamic.Headers
	}
	tests := []struct {
		name    string
		args    args
		want    *Header
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

			got, err := NewHeader(tt.args.next, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
