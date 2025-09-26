package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/httplib"
)

func TestMatchBodyFields_1(t *testing.T) {
	type args struct {
		ctx context.Context
		req *httplib.BeegoHTTPRequest
	}
	tests := []struct {
		name string
		sc   *SimpleCondition
		args args
		want bool
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

			if got := tt.sc.matchBodyFields(tt.args.ctx, tt.args.req); got != tt.want {
				t.Errorf("matchBodyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
