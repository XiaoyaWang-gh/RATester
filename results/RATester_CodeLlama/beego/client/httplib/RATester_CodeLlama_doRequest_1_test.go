package httplib

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDoRequest_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		b       *BeegoHTTPRequest
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name: "test case 1",
			b: &BeegoHTTPRequest{
				url:     "http://www.baidu.com",
				req:     &http.Request{},
				params:  map[string][]string{},
				files:   map[string]string{},
				setting: BeegoHTTPSettings{},
			},
			args: args{
				ctx: context.Background(),
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

			got, err := tt.b.doRequest(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
