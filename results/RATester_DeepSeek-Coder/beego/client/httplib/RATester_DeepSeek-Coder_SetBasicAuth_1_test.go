package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetBasicAuth_1(t *testing.T) {
	type args struct {
		username string
		password string
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

			if got := tt.b.SetBasicAuth(tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBasicAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
