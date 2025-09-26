package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReqBody_1(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		b    *BeegoHTTPRequest
		args args
		want *BeegoHTTPRequest
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

			if got := tt.b.reqBody(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeegoHTTPRequest.reqBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
