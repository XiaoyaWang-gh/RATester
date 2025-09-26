package try

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestDoRequest_1(t *testing.T) {
	type args struct {
		action     timedAction
		timeout    time.Duration
		request    *http.Request
		transport  http.RoundTripper
		conditions []ResponseCondition
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				action:     nil,
				timeout:    time.Duration(0),
				request:    nil,
				transport:  nil,
				conditions: nil,
			},
			want:    &http.Response{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := doRequest(tt.args.action, tt.args.timeout, tt.args.request, tt.args.transport, tt.args.conditions...)
			if (err != nil) != tt.wantErr {
				t.Errorf("doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
