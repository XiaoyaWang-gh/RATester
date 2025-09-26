package httplib

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetResponse_1(t *testing.T) {
	type args struct {
		b *BeegoHTTPRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			"case 1",
			args{
				b: &BeegoHTTPRequest{
					url: "http://www.baidu.com",
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.args.b.getResponse()
			if (err != nil) != tt.wantErr {
				t.Errorf("getResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
