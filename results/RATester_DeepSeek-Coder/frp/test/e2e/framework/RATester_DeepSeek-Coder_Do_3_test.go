package framework

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestDo_3(t *testing.T) {
	type fields struct {
		req         *request.Request
		f           *Framework
		expectResp  []byte
		expectError bool
		explain     []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    *request.Response
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

			e := &RequestExpect{
				req:         tt.fields.req,
				f:           tt.fields.f,
				expectResp:  tt.fields.expectResp,
				expectError: tt.fields.expectError,
				explain:     tt.fields.explain,
			}
			got, err := e.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestExpect.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestExpect.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
