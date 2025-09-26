package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOptions_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		r       *Request
		args    args
		want    *Response
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

			got, err := tt.r.Options(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.Options() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.Options() = %v, want %v", got, tt.want)
			}
		})
	}
}
