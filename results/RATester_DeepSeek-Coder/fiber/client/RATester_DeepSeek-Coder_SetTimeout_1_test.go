package client

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSetTimeout_1(t *testing.T) {
	type fields struct {
		timeout time.Duration
	}
	type args struct {
		t time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Request
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

			r := &Request{
				timeout: tt.fields.timeout,
			}
			if got := r.SetTimeout(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.SetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
