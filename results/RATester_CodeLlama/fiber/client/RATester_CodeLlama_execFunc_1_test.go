package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExecFunc_1(t *testing.T) {
	type args struct {
		c *core
	}
	tests := []struct {
		name    string
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				c: &core{
					client: &Client{},
					req: &Request{
						client: &Client{},
					},
				},
			},
			want:    &Response{},
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

			got, err := tt.args.c.execFunc()
			if (err != nil) != tt.wantErr {
				t.Errorf("execFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("execFunc() got = %v, want %v", got, tt.want)
			}
		})
	}
}
