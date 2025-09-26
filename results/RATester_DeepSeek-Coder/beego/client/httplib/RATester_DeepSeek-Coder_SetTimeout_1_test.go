package httplib

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSetTimeout_1(t *testing.T) {
	type args struct {
		connectTimeout   time.Duration
		readWriteTimeout time.Duration
	}
	tests := []struct {
		name string
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

			b := &BeegoHTTPRequest{}
			if got := b.SetTimeout(tt.args.connectTimeout, tt.args.readWriteTimeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeegoHTTPRequest.SetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
